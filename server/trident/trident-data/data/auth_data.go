package data

import (
	nc "nvm-gocore"
	"time"
)

type TimeSpan struct {
	Start time.Time
	End   time.Time
}

type AuthSessionData struct {
	Status          int
	Token           nc.Option[string]
	User            nc.Option[UserData]
	SessionDuration nc.Option[TimeSpan]
	Expired         nc.Option[time.Time]
}

type AuthData struct {
	Username   string
	Password   string
	Scope      string
	Persistent bool
}
