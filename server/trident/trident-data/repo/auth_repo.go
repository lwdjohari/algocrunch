package tridentdata

import (
	nc "nvm-gocore"
	td "trident-data"
)

type AuthRepo struct {
}

func NewAuthRepo() AuthRepo {
	return AuthRepo{}
}

func (ar *AuthRepo) DoAuth(
	username string,
	password string,
	scope string, context nc.Option[td.ExecutionContext],
	result chan nc.Result[bool]) {

	result <- nc.NewResult[bool](true, nil)
}

func (ar *AuthRepo) DoLogout(
	token string, scope string,
	context nc.Option[td.ExecutionContext],
	result chan nc.Result[bool]) {

	result <- nc.NewResult[bool](true, nil)
}
