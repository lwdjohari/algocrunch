package data

import "time"

type UserStatus uint32

const (
	UserStatus_INACTIVE    = 0
	UserStatus_ACTIVE      = 1
	UserStatus_INVITED     = 2
	UserStatus_UNCONFIRMED = 4
	UserStatus_UNVERIFIED  = 8
	UserStatus_SUSPENDED   = 16
	UserStatus_LOCKED      = 32
	UserStatus_DELETED     = 64
)

type IdentityType uint16

const (
	IdentityType_UNKNOWN        = 0
	IdentityType_USERNAME       = 1
	IdentityType_EMAIL          = 2
	IdentityType_PHONE          = 4
	IdentityType_SERVICE_SCOPED = 8
)

type IdentityProviderType uint32

const (
	IdentityProviderType_UNKNOWN   = 0
	IdentityProviderType_INTERNAL  = 1
	IdentityProviderType_EXT_AUTH  = 2
	IdentityProviderType_GOOGLE    = 4
	IdentityProviderType_FACEBOOK  = 8
	IdentityProviderType_TWITTER   = 16
	IdentityProviderType_GITHUB    = 32
	IdentityProviderType_LINKEDIN  = 64
	IdentityProviderType_MICROSOFT = 128
	IdentityProviderType_YAHOO     = 256
	IdentityProviderType_APPLE     = 512
	IdentityProviderType_TWITCH    = 1024
	IdentityProviderType_AMAZON    = 2048
)

type UserData struct {
	UserId           uint64
	Username         string
	Email            string
	Phone            string
	Password         string
	IdentityType     IdentityType
	IdentityProvider IdentityProviderType
	Flags            uint64
	Status           UserStatus
	CreateOn         time.Time
	CreateBy         uint64
	UpdateOn         time.Time
	UpdateBy         uint64
}

type UserProfileData struct {
	UserId              uint64
	Name                string
	FirstName           string
	LastName            string
	IsNotHavingLastName bool
	Departement         string
	Company             string
	CreateOn            time.Time
	CreateBy            uint64
	UpdateOn            time.Time
	UpdateBy            uint64
}

type UserExtendedInfo struct {
	UserExtendedInfoId uint64
	UserId             uint64
	Key                string
	Caption            string
	Value              string
	Type               int32
	CreateOn           time.Time
	CreateBy           uint64
	UpdateOn           time.Time
	UpdateBy           uint64
}

type UserAvatar struct {
	AvatarId     uint64
	UserId       uint64
	Caption      string
	Url          string
	ThumbnailUrl string
	Flags        uint64
	CreateOn     time.Time
	CreateBy     uint64
	UpdateOn     time.Time
	UpdateBy     uint64
}
