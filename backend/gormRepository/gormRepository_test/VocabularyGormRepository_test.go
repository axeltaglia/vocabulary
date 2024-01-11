package gormRepository_test

import (
	"testing"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/gormRepository"
	"vocabulary/gormRepository/VocabularyGormRepository"
)

// docker run --name pg_vocabulary_test_ctn -e POSTGRES_USER=vocabulary_test -e POSTGRES_PASSWORD=vocabulary_test -e POSTGRES_DB=vocabulary_test -e PGPORT=5436 -p 5436:5436 -d postgres
func TestVocabularyGormRepository(t *testing.T) {
	t.Run("CreateVocabulary", TestCreateVocabulary)
}
func TestCreateVocabulary(t *testing.T) {
	db, _ := gormRepository.ConnectToDbWithMaxAttempts(gormRepository.DbConfig{
		Host:     "localhost",
		Port:     "5436",
		User:     "vocabulary_test",
		Password: "vocabulary_test",
		Database: "vocabulary_test",
	}, 30)
	repository := VocabularyGormRepository.New(db)
	createdVocabulary := repository.CreateVocabulary(VocabularyEntity.Vocabulary{
		Words:        PString("Words"),
		Translation:  PString("Translation"),
		UsedInPhrase: PString("UsedInPhrase"),
		Explanation:  PString("Explanation"),
	})
	if createdVocabulary.Id == nil {
		t.Errorf("Vocabulary Id is expected")
	}

	if createdVocabulary.CreatedAt == nil {
		t.Errorf("Vocabulary CreatedAt date is expected")
	}

	if createdVocabulary.UpdatedAt == nil {
		t.Errorf("Vocabulary UpdatedAt date is expected")
	}

	if createdVocabulary.Words == nil || *createdVocabulary.Words != "Words" {
		t.Errorf("Vocabulary Word field is expected")
	}
}

func PString(s string) *string {
	return &s
}
