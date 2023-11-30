package entities

import "vocabulary/entities/VocabularyEntity"

type RepositoryFactory interface {
	BeginTransaction()
	CommitTransaction()
	RollbackTransaction()
	GetVocabularyRepository() VocabularyEntity.VocabularyRepository
}
