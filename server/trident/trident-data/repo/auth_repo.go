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
	serviceRepo ServiceRepo
}

func NewAuthRepo() AuthRepo {
	return AuthRepo{
		serviceRepo: *NewService(true),
	}
}

func (ar *AuthRepo) DoAuth(
	data dt.AuthData,
	context nc.Option[ns.ExecutionContext],
	result chan nc.Result[*dt.AuthSessionData]) {

	if context.IsNone() {
		result <- nc.NewResult[*dt.AuthSessionData](nil, errors.New("no excution context found"))
	}

	ctx := context.Unwrap()

	if ctx.IsDbContext() {
		ctx.DbContext().Unwrap().DB.Query("")
	} else {

	}
	result <- nc.NewResult[*dt.AuthSessionData](nil, nil)
}

func (ar *AuthRepo) DoLogout(
	token string, scope string,
	context nc.Option[ns.ExecutionContext],
	result chan nc.Result[bool]) {

	result <- nc.NewResult[bool](true, nil)
}
