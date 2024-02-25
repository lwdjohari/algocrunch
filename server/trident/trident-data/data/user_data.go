package data

type UserStatus uint32

const (
	UserStatus_DISABLE = 0
	UserStatus_ACTIVE  = 1
	UserStatus_LOCKED  = 2
)

type UserData struct {
	UserId   uint64
	Username string
	Email    string
	Phone    string
	Status   UserStatus
}
