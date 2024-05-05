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
	query := `insert into vocabularies (words, translation, used_in_phrase, explanation, created_at, updated_at)
				values ($1, $2, $3, $4, $5, $6) RETURNING id`
	now := time.Now()
	row := o.tx.QueryRow(query, *vocabulary.Words, *vocabulary.Translation, *vocabulary.UsedInPhrase, *vocabulary.Explanation, now, now)
	if row.Err() != nil {
		return nil, row.Err()
	}
	lastInsertId := 0
	if err := row.Scan(&lastInsertId); err != nil {
		logger.GetLogger().LogInfo("Error executing SQL statement: " + err.Error())
		return nil, err
	}

	uintId := uint(lastInsertId)
	vocabulary.Id = &uintId

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
	query := "insert into vocabulary_categories (vocabulary_id, category_id) values ($1, $2)"
	row, err := o.tx.Query(query, *vocabulary.Id, *category.Id)
	if err != nil {
		logger.GetLogger().LogInfo("Error executing SQL statement: " + err.Error())
		return nil, err
	}

	_ = row.Close()

	vocabulary.Categories = append(vocabulary.Categories, *category)
	return vocabulary, nil
}

func (o *SqlRepository) CreateCategoryIfNotExist(categoryName string) (*VocabularyEntity.Category, error) {
	var category VocabularyEntity.Category

	query := "select * from categories where name=$1"
	row, err := o.tx.Query(query, categoryName)
	if err != nil {
		logger.GetLogger().LogInfo("Error executing SQL statement: " + err.Error())
		return nil, err
	}

	rowExist := row.Next()
	if rowExist {
		if err := row.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
			logger.GetLogger().LogError("Can't decode object", err)
			return nil, err
		}
		row.Close()
		return &category, nil
	} else {
		now := time.Now()
		query := "insert into categories (name, created_at, updated_at) values ($1, $2, $3) RETURNING id"
		row, err := o.tx.Query(query, categoryName, now, now)
		if err != nil {
			logger.GetLogger().LogInfo("Error executing SQL statement: " + err.Error())
			return nil, err
		}

		var lastInsertId uint
		row.Next()
		if err := row.Scan(&lastInsertId); err != nil {
			logger.GetLogger().LogInfo("Error executing SQL statement: " + err.Error())
			return nil, err
		}
		row.Close()

		category.CreatedAt = &now
		category.UpdatedAt = &now
		category.Name = &categoryName
		category.Id = &lastInsertId

		return &category, nil
	}
}

func (o *SqlRepository) DeleteVocabularyById(u uint) error {
	//TODO implement me
	panic("implement me")
}

func NewSqlRepository(tx *sql.Tx) *SqlRepository {
	return &SqlRepository{tx: tx}
}
