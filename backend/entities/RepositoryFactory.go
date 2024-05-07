package entities

import (
	"vocabulary/entities/VocabularyEntity"
)

type TxRepositoryHandler interface {
	GetTxRepositoryFactory() TxRepositoryFactory
}

type TxRepositoryFactory interface {
	CommitTransaction()
	RollbackTransaction()
	TransactionError() *string
	GetVocabularyRepository() VocabularyEntity.VocabularyRepository

	// add more repositories getters, like "GetUserRepository"
}
