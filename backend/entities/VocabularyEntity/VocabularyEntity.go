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

type VocabularyRepository interface {
	CreateVocabulary(vocabulary Vocabulary) Vocabulary
	GetAllVocabulariesWithCategories() []Vocabulary
	FindVocabularyById(id uint) Vocabulary
}
type Entity struct {
	Repository VocabularyRepository
}

func (o Entity) Create(vocabulary Vocabulary) Vocabulary {
	newVocabulary := o.Repository.CreateVocabulary(Vocabulary{
		Words:        vocabulary.Words,
		Translation:  vocabulary.Translation,
		UsedInPhrase: vocabulary.UsedInPhrase,
		Explanation:  vocabulary.Explanation,
	})

	return newVocabulary
}

func (o Entity) GetAllVocabulariesWithCategories() []Vocabulary {
	return o.Repository.GetAllVocabulariesWithCategories()
}

func (o Entity) GetVocabulary(id uint) Vocabulary {
	return o.Repository.FindVocabularyById(id)
}

func New(repository VocabularyRepository) Entity {
	return Entity{
		Repository: repository,
	}
}
