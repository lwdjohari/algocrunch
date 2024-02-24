package tridentcore

type TridentServiceType int32

const (
	TridentService_UNKNOWN TridentServiceType = 0
	TridentService_AUTH    TridentServiceType = 1
	TridentService_USER    TridentServiceType = 2
)

type ITridentService interface {
	ServiceType() TridentServiceType
}
