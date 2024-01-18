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

func (o Entity) Create(vocabulary *Vocabulary) (*Vocabulary, error) {
	newVocabulary, err := o.Repository.CreateVocabulary(vocabulary)
	if err != nil {
		logger.LogInfo("Vocabulary couldn't be created")
		return nil, err
	}
	return newVocabulary, err
}

func (o Entity) CreateWithCategories(vocabulary *Vocabulary, categories []string) (*Vocabulary, error) {
	newVocabulary, err := o.Repository.CreateVocabularyWithCategories(vocabulary, categories)
	if err != nil {
		logger.LogInfo("Vocabulary couldn't be created")
		return nil, err
	}
	return newVocabulary, err
}

func (o Entity) GetAllVocabulariesWithCategories() ([]Vocabulary, error) {
	vocabularies, err := o.Repository.GetAllVocabulariesWithCategories()
	if err != nil {
		logger.LogError("GetAllVocabulariesWithCategories has failed", err)
		return nil, err
	}
	return vocabularies, nil
}

func (o Entity) GetVocabulary(id uint) (*Vocabulary, error) {
	vocabulary, err := o.Repository.FindVocabularyById(id)
	if err != nil {
		logger.LogError("GetVocabulary has failed", err)
		return nil, err
	}
	return vocabulary, nil
}

func (o Entity) GetCategoriesFromVocabulary(id uint) ([]Category, error) {
	vocabulary, err := o.Repository.FindVocabularyById(id)
	if err != nil {
		logger.LogError("GetVocabulary has failed", err)
		return nil, err
	}
	return vocabulary.Categories, err
}

func (o Entity) GetAllCategories() []Category {
	categories := o.Repository.FindCategories()
	return categories
}

func (o Entity) Update(vocabulary *Vocabulary) (*Vocabulary, error) {
	updatedVocabulary, _ := o.Repository.UpdateVocabulary(vocabulary)
	return updatedVocabulary, nil
}

func (o Entity) Delete(id uint) error {
	if err := o.Repository.DeleteVocabularyById(id); err != nil {
		return err
	}
	return nil
}

func (o Entity) UpdateWithCategories(vocabulary *Vocabulary, categories []string) (*Vocabulary, error) {
	updatedVocabulary, _ := o.Repository.UpdateVocabularyWithCategories(vocabulary, categories)
	return updatedVocabulary, nil
}

func New(repository VocabularyRepository) Entity {
	return Entity{
		Repository: repository,
	}
}
