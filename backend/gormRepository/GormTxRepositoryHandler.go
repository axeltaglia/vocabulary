package gormRepository

import (
	"github.com/jinzhu/gorm"
	"vocabulary/entities"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/gormRepository/VocabularyGormRepository"
)

type GormTxRepositoryHandler struct {
	db *gorm.DB
}

func (o *GormTxRepositoryHandler) GetTxRepositoryFactory() entities.TxRepositoryFactory {
	newTx := o.db.Begin()
	return &GormTxRepositoryFactory{tx: newTx}
}

func NewGormTxRepositoryHandler(db *gorm.DB) *GormTxRepositoryHandler {
	return &GormTxRepositoryHandler{db: db}
}

type GormTxRepositoryFactory struct {
	tx *gorm.DB
}

func (o *GormTxRepositoryFactory) CommitTransaction() {
	o.tx.Commit()
}

func (o *GormTxRepositoryFactory) TransactionError() *string {
	if o.tx.Error != nil {
		errMsg := o.tx.Error.Error()
		return &errMsg
	}
	return nil
}

func (o *GormTxRepositoryFactory) RollbackTransaction() {
	o.tx.Rollback()
}

func (o *GormTxRepositoryFactory) CreateVocabularyRepository() VocabularyEntity.VocabularyRepository {
	return VocabularyGormRepository.New(o.tx)
}
