package sqlRepository

import (
	"database/sql"
	"time"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/logger"
)

type SqlRepository struct {
	tx *sql.Tx
}

func (o *SqlRepository) CreateVocabulary(vocabulary *VocabularyEntity.Vocabulary) (*VocabularyEntity.Vocabulary, error) {
	query := `insert into vocabulary (words, translation, used_in_phrase, explanation, created_at, updated_at)
				values ($1, $2, $3, $4, $5)`
	now := time.Now().UTC().String()
	var lastInsertedId uint
	if err := o.tx.QueryRow(query, *vocabulary.Words, *vocabulary.Translation, *vocabulary.Explanation, now, now).Scan(&lastInsertedId); err != nil {
		logger.GetLogger().LogError(err.Error(), err)
		return nil, err
	}

	vocabulary.Id = &lastInsertedId

	return vocabulary, nil
}

func (o *SqlRepository) CreateVocabularyWithCategories(vocabulary *VocabularyEntity.Vocabulary, strings []string) (*VocabularyEntity.Vocabulary, error) {
	//TODO implement me
	panic("implement me")
}

func (o *SqlRepository) GetAllVocabulariesWithCategories() ([]VocabularyEntity.Vocabulary, error) {
	//TODO implement me
	panic("implement me")
}

func (o *SqlRepository) FindVocabularyById(id uint) (*VocabularyEntity.Vocabulary, error) {
	//TODO implement me
	panic("implement me")
}

func (o *SqlRepository) FindCategories() []VocabularyEntity.Category {
	//TODO implement me
	panic("implement me")
}

func (o *SqlRepository) FindCategoriesByVocabularyId(id uint) ([]VocabularyEntity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (o *SqlRepository) UpdateVocabulary(vocabulary *VocabularyEntity.Vocabulary) (*VocabularyEntity.Vocabulary, error) {
	//TODO implement me
	panic("implement me")
}

func (o *SqlRepository) DisassociateCategoriesFromVocabulary(vocabulary *VocabularyEntity.Vocabulary) error {
	//TODO implement me
	panic("implement me")
}

func (o *SqlRepository) AssociateCategoryToVocabulary(vocabulary *VocabularyEntity.Vocabulary, category *VocabularyEntity.Category) (*VocabularyEntity.Vocabulary, error) {
	//TODO implement me
	panic("implement me")
}

func (o *SqlRepository) CreateCategoryIfNotExist(categoryName string) (*VocabularyEntity.Category, error) {
	//TODO implement me
	panic("implement me")
}

func (o *SqlRepository) DeleteVocabularyById(u uint) error {
	//TODO implement me
	panic("implement me")
}

func NewSqlRepository(tx *sql.Tx) *SqlRepository {
	return &SqlRepository{tx: tx}
}
