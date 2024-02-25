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

type AuthRepo struct {
}

func NewAuthRepo() AuthRepo {
	return AuthRepo{}
}

func (ar *AuthRepo) DoAuth(
	username string,
	password string,
	scope string,
	context nc.Option[ns.ExecutionContext],
	result chan nc.Result[*dt.AuthData]) {

	if context.IsNone() {
		result <- nc.NewResult[*dt.AuthData](nil, errors.New("no excution context found"))
	}

	ctx := context.Unwrap()

	if ctx.IsDbContext() {
		ctx.DbContext().Unwrap().DB.Query("")
	} else {

	}
	result <- nc.NewResult[*dt.AuthData](nil, nil)
}

func (ar *AuthRepo) DoLogout(
	token string, scope string,
	context nc.Option[ns.ExecutionContext],
	result chan nc.Result[bool]) {

	result <- nc.NewResult[bool](true, nil)
}
