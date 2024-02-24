package tridentcore

import (
	"fmt"
	nc "nvm-gocore"
	tc "trident-core"
	td "trident-data"
)

type AuthService struct {
	service_type tc.TridentServiceType
}

func NewAuthService() AuthService {
	return AuthService{
		service_type: tc.TridentService_AUTH,
	}
}

func (as *AuthService) ServiceType() tc.TridentServiceType {
	return as.service_type
}

func (as *AuthService) AuthAsync(
	username string,
	password string,
	scope string,
	contex nc.Option[td.ExecutionContext],
	result chan<- nc.Result[bool]) {
	result <- nc.NewResult[bool](false, nil)
}

func (as *AuthService) ValidateTokenAsync(
	token string,
	contex nc.Option[td.ExecutionContext],
	result chan<- nc.Result[bool],
) {
	result <- nc.NewResult[bool](false, nil)
}

func (as *AuthService) Logout(
	token string,
	contex nc.Option[td.ExecutionContext],
	result chan<- nc.Result[bool],
) {
	result <- nc.NewResult[bool](false, nil)
}

func Example() {
	authService := NewAuthService()

	channel := nc.MakeChannel[nc.Result[bool]]()
	authService.AuthAsync("", "", "", nc.None[td.ExecutionContext](), channel)

	result := nc.GetFromChannel[nc.Result[bool]]()

	if result.IsOk() {
		fmt.Println(result.Value)
	} else {
		fmt.Println(result.Err)
	}

}
