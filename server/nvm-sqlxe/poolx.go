package nvmsqlxe

import (
	_ "database/sql"
	"errors"
	"math"
	nc "nvm-gocore"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type PoolOpenStrategies uint32

const (
	PoolOpenStrategies_UNKNOWN      = 0
	PoolOpenStrategies_OPEN_ALL     = 1
	PoolOpenStrategies_ON_REQUEST   = 2
	PoolOpenStrategies_OPEN_MINIMAL = 3
)

type ConnectionInfoX interface {
	DbBackendType() DbBackendType
	GetConnectionString() string
}

type PoolX struct {
	mu          sync.Mutex
	cond        *sync.Cond
	connections []*sqlx.DB
	leased      map[*sqlx.DB]bool
	unleased    map[*sqlx.DB]bool

	maxConnectionNumber uint16
	connectionInfo      ConnectionInfoX
	connOpenMinimal     uint16
	connOpened          uint16
	connUnopened        uint16
	openStrategy        PoolOpenStrategies
	isRun               bool
}

func NewPoolX(connInfo ConnectionInfoX, maxConnectionNumber uint16) *PoolX {
	px := &PoolX{
		connectionInfo:      connInfo,
		maxConnectionNumber: maxConnectionNumber,
		connOpenMinimal:     0,
		connOpened:          0,
		connUnopened:        0,
		openStrategy:        PoolOpenStrategies_UNKNOWN,
		connections:         make([]*sqlx.DB, 0, maxConnectionNumber),
		unleased:            make(map[*sqlx.DB]bool),
		leased:              make(map[*sqlx.DB]bool),
		isRun:               false,
	}

	px.cond = sync.NewCond(&px.mu)
	return px
}

func (px *PoolX) getDbBackendTypeAsString() nc.Option[string] {
	switch px.connectionInfo.DbBackendType() {
	case DbBackendType_PG:
		return nc.Some[string]("postgress")
	case DbBackendType_MYSQL:
		return nc.Some[string]("mysql")
	case DbBackendType_SQLITE:
		return nc.Some[string]("sqlite")
	default:
		return nc.None[string]()
	}
}

func (px *PoolX) openConnection() nc.Result[*sqlx.DB] {
	backend := px.getDbBackendTypeAsString()
	if backend.IsNone() {
		return nc.NewResult[*sqlx.DB](nil, errors.New("unsupported DB Backend"))
	}

	db, err := sqlx.Open(backend.Unwrap(), px.connectionInfo.GetConnectionString())
	if err != nil {
		return nc.NewResult[*sqlx.DB](nil, err)
	}

	return nc.NewResult[*sqlx.DB](db, nil)
}

func (px *PoolX) Start(
	openStrategy PoolOpenStrategies,
	minOpenConnection uint16,
	maxErrorBreaker uint16,
) nc.Result[uint16] {
	px.openStrategy = openStrategy
	px.connOpenMinimal = minOpenConnection

	if maxErrorBreaker == 0 {
		ceilResult := math.Ceil(float64(px.maxConnectionNumber) * 0.3)
		if ceilResult < 1 {
			maxErrorBreaker = 1
		} else {
			maxErrorBreaker = uint16(ceilResult)
		}
	}

	px.mu.Lock()
	defer px.mu.Unlock()

	numConnSuccess := 0
	numConnFailed := 0
	var err error

	if openStrategy == PoolOpenStrategies_OPEN_ALL {
		for i := 0; i < int(px.maxConnectionNumber); i++ {
			res := px.openConnection()
			if res.IsOk() {
				_ = append(px.connections, res.Value)
				px.unleased[res.Value] = true
				numConnSuccess++
			} else {
				numConnFailed++
				err = res.Err
				if numConnFailed == int(maxErrorBreaker) {
					break
				}
			}

		}

		px.isRun = true
		return nc.NewResult[uint16](uint16(numConnSuccess), err)
	} else if openStrategy == PoolOpenStrategies_OPEN_MINIMAL {
		if minOpenConnection == 0 || minOpenConnection > px.maxConnectionNumber {
			return nc.NewResult[uint16](0, errors.New("minimal connection zero or larger than max connections number"))
		}

		if maxErrorBreaker > minOpenConnection {
			maxErrorBreaker = minOpenConnection
		}

		for i := 0; i < int(minOpenConnection); i++ {
			res := px.openConnection()
			if res.IsOk() {
				_ = append(px.connections, res.Value)
				px.unleased[res.Value] = true
				numConnSuccess++
			} else {
				numConnFailed++
				err = res.Err
				if numConnFailed == int(maxErrorBreaker) {
					break
				}
			}

		}

		px.isRun = true
		return nc.NewResult[uint16](uint16(numConnSuccess), err)

	} else if openStrategy == PoolOpenStrategies_ON_REQUEST {
		px.isRun = true
		return nc.NewResult[uint16](0, nil)
	} else {
		return nc.NewResult[uint16](0, errors.New("poolOpenStrategies_UNKNOWN detected, please choose one open strategy"))
	}

}

func (px *PoolX) waitForConnection(timeout time.Duration) bool {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			return false // Timeout occurred
		default:
			px.cond.Wait() // Wait for a signal that a connection has been returned.
			if (len(px.connections) - len(px.leased)) > 0 {
				return true // A connection has become available.
			}
		}
	}
}

func (px *PoolX) leaseOut() nc.Result[*sqlx.DB] {
	k, v, f := nc.MapGetOneOf[*sqlx.DB, bool](px.unleased)
	if f {
		res := nc.NewResult[*sqlx.DB](k, nil)
		px.leased[k] = v
		delete(px.unleased, k)
		return res
	} else {
		return nc.NewResult[*sqlx.DB](nil, errors.New("db conn leased failed"))
	}
}

// Leased DB Connection  when 0 will wait forever until get connection.
// After using must return the DB Connection to pool, otherwise you will end having no DB Connection.
func (px *PoolX) Get(timeoutInSecond uint) nc.Result[*sqlx.DB] {
	px.mu.Lock()
	defer px.mu.Unlock()

	if timeoutInSecond == 0 {
		px.cond.Wait() // Wait for a signal that a connection has been returned.
		if (len(px.connections) - len(px.leased)) > 0 {
			return px.leaseOut()
		}
	} else {
		if len(px.unleased) > 0 {
			return px.leaseOut()
		} else {
			// Wait for a connection to be returned or for the timeout.
			if !px.waitForConnection(time.Second * time.Duration(timeoutInSecond)) {
				return nc.NewResult[*sqlx.DB](nil, errors.New("timed out waiting for a connection"))
			}

			return px.leaseOut()
		}
	}

	return nc.NewResult[*sqlx.DB](nil, errors.New("Get DB Connection failed"))
}

func (px *PoolX) Return(db *sqlx.DB) bool {

	px.mu.Lock()
	defer px.mu.Unlock()

	leased, exists := px.leased[db]

	if exists {
		delete(px.leased, db)    // Remove from leased
		px.unleased[db] = leased // Add back to unleased
		px.cond.Signal()         // Signal that a connection has been returned.
		return true
	} else {
		return false
	}

}

func (px *PoolX) ConnectionAvailable() int {
	px.mu.Lock()
	defer px.mu.Unlock()
	return len(px.connections) - len(px.leased)
}

func (px *PoolX) ConnectionOpened() int {
	px.mu.Lock()
	defer px.mu.Unlock()
	return len(px.connections)
}

func (px *PoolX) ConnectionLeased() int {
	px.mu.Lock()
	defer px.mu.Unlock()
	return len(px.leased)
}

func (px *PoolX) MinConnectionNumber() uint16 {
	return px.connOpenMinimal
}

func (px *PoolX) MaxConnectionNumber() uint16 {
	return px.maxConnectionNumber
}

func (px *PoolX) OpenStrategy() PoolOpenStrategies {
	return px.openStrategy
}

func (px *PoolX) Stop() {
	px.mu.Lock()
	defer px.mu.Unlock()

	px.isRun = false

	for i := 0; i < len(px.connections); i++ {
		px.connections[i].Close()
	}

	clear(px.connections)
	clear(px.leased)

}

func (px *PoolX) DbBackendType() DbBackendType {
	return px.connectionInfo.DbBackendType()
}
