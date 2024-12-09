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
	CreateVocabularyRepository() VocabularyEntity.VocabularyRepository

	// add more repositories getters, like "CreateUserRepository" that creates a "UserEntity.UserRepository"
}
