package sqlRepository

import (
	"database/sql"
	"vocabulary/entities"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/logger"
)

type SqlTxRepositoryHandler struct {
	db *sql.DB
}

func NewSqlTxRepositoryHandler(db *sql.DB) SqlTxRepositoryHandler {
	return SqlTxRepositoryHandler{
		db: db,
	}
}

func (o *SqlTxRepositoryHandler) GetTxRepositoryFactory() entities.TxRepositoryFactory {
	tx, err := o.db.Begin()
	if err != nil {
		logger.GetLogger().LogError("couldn't start a transaction", err)
	}
	return &SqlTxRepositoryFactory{
		tx: tx,
	}
}

type SqlTxRepositoryFactory struct {
	tx *sql.Tx
}

func (o *SqlTxRepositoryFactory) CommitTransaction() {
	if err := o.tx.Commit(); err != nil {
		logger.GetLogger().LogInfo("couldn't commit the transaction")
	}
}

func (o *SqlTxRepositoryFactory) RollbackTransaction() {
	if err := o.tx.Rollback(); err != nil {
		logger.GetLogger().LogInfo("couln't rollback the transaction")
	}
}

func (o *SqlTxRepositoryFactory) TransactionError() *string {
	msg := ""
	return &msg
}

func (o *SqlTxRepositoryFactory) CreateVocabularyRepository() VocabularyEntity.VocabularyRepository {
	return NewSqlRepository(o.tx)
}
