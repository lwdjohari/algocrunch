package nvmsqlxe

import (
	nc "nvm-gocore"

	"github.com/jmoiron/sqlx"
)

type ExecutionContextType uint32

const (
	ExecutionContextType_UNKNOWN     = 0
	ExecutionContextType_DB_CONN     = 1
	ExecutionContextType_TRANSACTION = 2
)

type ExecutionContext struct {
	contextType ExecutionContextType
	dbContext   nc.Option[*sqlx.DB]
	transaction nc.Option[any]
}

func ExecutionContextFromDbContext(context nc.Option[*sqlx.DB]) *ExecutionContext {
	ec := &ExecutionContext{
		contextType: ExecutionContextType_DB_CONN,
		dbContext:   context,
		transaction: nc.None[any](),
	}

	return ec
}

func ExecutionContextFromTransaction(context nc.Option[any]) *ExecutionContext {
	ec := &ExecutionContext{
		contextType: ExecutionContextType_TRANSACTION,
		transaction: context,
		dbContext:   nc.None[*sqlx.DB](),
	}

	return ec
}

func (ec *ExecutionContext) ContextType() ExecutionContextType {
	return ec.contextType
}

func (ec *ExecutionContext) DbContext() nc.Option[*sqlx.DB] {
	return ec.dbContext
}

func (ec *ExecutionContext) DbTransaction() nc.Option[any] {
	return ec.transaction
}

func (ec *ExecutionContext) IsDbContext() bool {
	return ec.contextType == ExecutionContextType_DB_CONN
}

func (ec *ExecutionContext) IsTransactionContext() bool {
	return ec.contextType == ExecutionContextType_TRANSACTION
}
