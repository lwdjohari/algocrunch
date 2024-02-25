package nvmsqlxe

type DbBackendType uint32

const (
	DbBackendType_UNKNOWN DbBackendType = 0
	DbBackendType_PG      DbBackendType = 1
	DbBackendType_MYSQL   DbBackendType = 2
	DbBackendType_SQLITE  DbBackendType = 3
)
