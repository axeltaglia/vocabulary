package VocabularyEntity

import (
	"time"
	"vocabulary/logger"
)

type Vocabulary struct {
	Id           *uint
	Words        *string
	Translation  *string
	UsedInPhrase *string
	Explanation  *string
	Categories   []Category
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

type Category struct {
	Id           *uint        `json:"id"`
	Name         *string      `json:"name"`
	Vocabularies []Vocabulary `json:"vocabularies"`
	CreatedAt    *time.Time   `json:"createdAt"`
	UpdatedAt    *time.Time   `json:"updatedAt"`
}

type Entity struct {
	Repository VocabularyRepository
}

func (o Entity) CreateWithCategories(vocabulary *Vocabulary, categoryNames []string) (*Vocabulary, error) {
	newVocabulary, err := o.Repository.CreateVocabulary(vocabulary)
	if err != nil {
		return nil, err
	}

	for _, categoryName := range categoryNames {
		category, err := o.Repository.CreateCategoryIfNotExist(categoryName)
		if err != nil {
			return nil, err
		}

		newVocabulary, err = o.Repository.AssociateCategoryToVocabulary(newVocabulary, category)
		if err != nil {
			return nil, err
		}
	}

	return newVocabulary, nil
}

func (o Entity) GetAllVocabulariesWithCategories() ([]Vocabulary, error) {
	vocabularies, err := o.Repository.GetAllVocabulariesWithCategories()
	if err != nil {
		logger.GetLogger().LogError("GetAllVocabulariesWithCategories has failed", err)
		return nil, err
	}
	return vocabularies, nil
}

func (o Entity) GetVocabulary(id uint) (*Vocabulary, error) {
	return o.Repository.FindVocabularyById(id)
}

func (o Entity) GetCategoriesFromVocabulary(vocabularyId uint) ([]Category, error) {
	return o.Repository.FindCategoriesByVocabularyId(vocabularyId)
}

func (o Entity) GetAllCategories() []Category {
	categories := o.Repository.FindCategories()
	return categories
}

func (o Entity) Delete(id uint) error {
	if err := o.Repository.DeleteVocabularyById(id); err != nil {
		return err
	}
	return nil
}

func (o Entity) UpdateWithCategories(vocabulary *Vocabulary, categoryNames []string) (*Vocabulary, error) {
	updatedVocabulary, err := o.Repository.UpdateVocabulary(vocabulary)
	if err != nil {
		return nil, err
	}

	err = o.Repository.DisassociateCategoriesFromVocabulary(updatedVocabulary)
	if err != nil {
		return nil, err
	}

	for _, categoryName := range categoryNames {
		category, err := o.Repository.CreateCategoryIfNotExist(categoryName)
		if err != nil {
			return nil, err
		}

		updatedVocabulary, err = o.Repository.AssociateCategoryToVocabulary(updatedVocabulary, category)
		if err != nil {
			return nil, err
		}
	}

	return updatedVocabulary, nil
}

func New(repository VocabularyRepository) Entity {
	return Entity{
		Repository: repository,
	}
}
