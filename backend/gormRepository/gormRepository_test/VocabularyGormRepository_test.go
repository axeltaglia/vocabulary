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
	db := gormRepository.InitDb(gormRepository.DbConfig{
		Host:     "localhost",
		Port:     "5435",
		User:     "vocabulary",
		Password: "vocabulary",
		Database: "vocabulary",
	})
	repository := VocabularyGormRepository.New(db)
	createdVocabulary := repository.CreateVocabulary(VocabularyEntity.Vocabulary{
		Words:        PString("Words"),
		Translation:  PString("Translation"),
		UsedInPhrase: PString("UsedInPhrase"),
		Explanation:  PString("Explanation"),
	})
	if createdVocabulary.Id == nil {
		t.Errorf("A vocabulary Id is expected")
	}
}

func PString(s string) *string {
	return &s
}
