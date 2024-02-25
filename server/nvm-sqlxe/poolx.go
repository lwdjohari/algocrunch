package nvmsqlxe

import (
	_ "database/sql"
	"errors"
	nc "nvm-gocore"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type PoolOpenStrategies uint32

type ConnectionInfoX interface {
	DbBackendType() DbBackendType
	GetConnectionString() string
}

type PoolX struct {
	connections         *sqlx.DB // Connection, sqlx has already connection pool feature
	maxConnectionNumber uint16
	maxConnectionIdle   uint16
	maxConnLifeInSecond uint32
	connectionInfo      ConnectionInfoX

	isInitialize bool
}

func NewPoolX(
	connInfo ConnectionInfoX,
	maxConnectionNumber uint16,
	maxConnectionIdle uint16,
	maxConnLifeInSecond uint32,
) *PoolX {
	px := &PoolX{
		connectionInfo:      connInfo,
		maxConnectionNumber: maxConnectionNumber,
		maxConnectionIdle:   maxConnectionIdle,
		maxConnLifeInSecond: maxConnLifeInSecond,
		connections:         nil,
		isInitialize:        false,
	}

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

func (px *PoolX) Open() nc.Result[bool] {

	if px.isInitialize {
		return nc.NewResult[bool](true, nil)
	}

	backend := px.getDbBackendTypeAsString()
	if backend.IsNone() {
		return nc.NewResult[bool](false, errors.New("unsupported DB Backend"))
	}

	db, err := sqlx.Open(backend.Unwrap(), px.connectionInfo.GetConnectionString())
	if err != nil {
		return nc.NewResult[bool](false, err)
	}

	maxLifeTime := time.Duration(1000 * px.maxConnLifeInSecond)
	db.SetMaxOpenConns(int(px.maxConnectionNumber))
	db.SetMaxIdleConns(int(px.maxConnectionIdle))
	db.SetConnMaxLifetime(maxLifeTime)

	px.connections = db

	px.isInitialize = true
	return nc.NewResult[bool](true, nil)
}

func (px *PoolX) GetDB() *sqlx.DB {
	return px.connections
}

func (px *PoolX) GetDbContext() nc.Result[*ExecutionContext] {
	if px.connections == nil {
		return nc.NewResult[*ExecutionContext](nil, errors.New("db context null-reference"))
	} else {
		return nc.NewResult[*ExecutionContext](ExecutionContextFromDbContext(nc.Some[*sqlx.DB](px.connections), px.connectionInfo.DbBackendType()), nil)
	}
}

func (px *PoolX) GetTransactionContext() nc.Result[*ExecutionContext] {
	if px.connections == nil {
		return nc.NewResult[*ExecutionContext](nil, errors.New("db context null-reference"))
	} else {
		tx, err := px.connections.Beginx()
		if err != nil {
			return nc.NewResult[*ExecutionContext](nil, err)
		}
		return nc.NewResult[*ExecutionContext](ExecutionContextFromTransaction(nc.Some[*sqlx.Tx](tx), px.connectionInfo.DbBackendType()), nil)
	}
}

func (px *PoolX) Close() bool {
	if !px.isInitialize {
		return true
	}

	err := px.connections.Close()
	px.isInitialize = false

	if err != nil {
		return false
	} else {
		return true
	}

}

func (px *PoolX) DbBackendType() DbBackendType {
	return px.connectionInfo.DbBackendType()
}
