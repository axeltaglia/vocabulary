package VocabularyEntity

import (
	"time"
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
	return newVocabulary, err
}

func (o Entity) GetAllVocabulariesWithCategories() []Vocabulary {
	return o.Repository.GetAllVocabulariesWithCategories()
}

func (o Entity) GetVocabulary(id uint) Vocabulary {
	return o.Repository.FindVocabularyById(id)
}

func (o Entity) GetCategoriesFromVocabulary(vocabularyId uint) []Category {
	vocabulary := o.Repository.FindVocabularyById(vocabularyId)
	return vocabulary.Categories
}

func (o Entity) GetAllCategories() []Category {
	categories := o.Repository.FindCategories()
	return categories
}

func (o Entity) Update(vocabulary Vocabulary) Vocabulary {
	return o.Repository.UpdateVocabulary(vocabulary)
}

func (o Entity) UpdateWithCategories(vocabulary Vocabulary, categories []string) {
	o.Repository.UpdateVocabularyWithCategories(vocabulary, categories)
}

func New(repository VocabularyRepository) Entity {
	return Entity{
		Repository: repository,
	}
}
