package data

import (
	"time"
)

type ServiceStatus uint8

const (
	ServiceStatus_INACTIVE             = 0
	ServiceStatus_ACTIVE               = 1
	ServiceStatus_SUSPENDED            = 2
	ServiceStatus_WAITING_CONFIRMATION = 4
	ServiceStatus_WAITING_VERIFICATION = 8
	ServiceStatus_DELETED              = 16
)

type ServiceData struct {
	ServiceId uint64
	Key       string
	Name      string
	Status    ServiceStatus
	CreateOn  time.Time
	CreateBy  uint64
	UpdateOn  time.Time
	UpdateBy  uint64
}
