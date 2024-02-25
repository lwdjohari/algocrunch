package data

type AuthData struct {
	Status int
	Token  *string
	User   *UserData
}

func NewAuthData() AuthData {
	return AuthData{
		Status: 0,
		Token:  nil,
		User:   nil,
	}
}
