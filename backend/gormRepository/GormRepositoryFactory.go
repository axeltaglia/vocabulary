package gormRepository

import (
	"github.com/jinzhu/gorm"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/gormRepository/VocabularyGormRepository"
)

type Factory struct {
	db *gorm.DB
	tx *gorm.DB
}

func (o *Factory) BeginTransaction() {
	o.tx = o.db.Begin()
}

func (o *Factory) CommitTransaction() {
	o.checkTx()
	//o.tx.Commit()
}

func (o *Factory) RollbackTransaction() {
	o.checkTx()
	o.tx.Rollback()
}

func (o *Factory) GetVocabularyRepository() VocabularyEntity.VocabularyRepository {
	o.checkTx()
	return VocabularyGormRepository.New(o.tx)
}

func (o *Factory) checkTx() {
	if o.tx == nil {
		panic("First start the transaction")
	}
}

func NewFactory(db *gorm.DB) *Factory {
	return &Factory{
		db: db,
	}
}
