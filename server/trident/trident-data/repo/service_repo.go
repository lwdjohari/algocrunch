package repo

import (
	"errors"
	nc "nvm-gocore"
	ns "nvm-sqlxe"
	dt "trident-data/data"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type ServiceRepo struct {
	isUseCache bool
}

func NewService(cache bool) *ServiceRepo {
	sr := &ServiceRepo{
		isUseCache: cache,
	}
	return sr
}

func (sr *ServiceRepo) IsUseCache() bool {
	return sr.isUseCache
}

func (sr *ServiceRepo) GetScope(
	data dt.AuthData,
	context nc.Option[ns.ExecutionContext],
	invalidateCache bool,
	cacheDurationInSecond uint32,
	result chan nc.Result[*dt.ServiceData]) {

	if context.IsNone() {
		result <- nc.NewResult[*dt.ServiceData](nil, errors.New("no excution context found"))
	}

	ctx := context.Unwrap()

	if ctx.IsDbContext() {
		ctx.DbContext().Unwrap().DB.Query("")
	} else {

	}
	result <- nc.NewResult[*dt.ServiceData](nil, nil)
}
