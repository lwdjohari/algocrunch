package auth

import (
	nc "nvm-gocore"
	ns "nvm-sqlxe"
	tc "trident-core"
	tr "trident-data/repo"
)

type AuthService struct {
	service_type tc.TridentServiceType
	authRepo     tr.AuthRepo
}

func NewAuthService() AuthService {
	return AuthService{
		service_type: tc.TridentService_AUTH,
		authRepo:     tr.NewAuthRepo(),
	}
}

func (as *AuthService) ServiceType() tc.TridentServiceType {
	return as.service_type
}

func (as *AuthService) AuthAsync(
	username string,
	password string,
	scope string,
	contex nc.Option[ns.ExecutionContext],
	result chan<- nc.Result[bool]) {
	result <- nc.NewResult[bool](false, nil)

}

func (as *AuthService) ValidateTokenAsync(
	token string,
	contex nc.Option[ns.ExecutionContext],
	result chan<- nc.Result[bool],
) {
	result <- nc.NewResult[bool](false, nil)
}

func (as *AuthService) Logout(
	token string,
	contex nc.Option[ns.ExecutionContext],
	result chan<- nc.Result[bool],
) {
	result <- nc.NewResult[bool](false, nil)
}
