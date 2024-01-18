package VocabularyEntity

type VocabularyRepository interface {
	CreateVocabulary(vocabulary *Vocabulary) (*Vocabulary, error)
	CreateVocabularyWithCategories(*Vocabulary, []string) (*Vocabulary, error)
	GetAllVocabulariesWithCategories() ([]Vocabulary, error)
	FindVocabularyById(id uint) (*Vocabulary, error)
	FindCategories() []Category
	UpdateVocabulary(vocabulary *Vocabulary) (*Vocabulary, error)
	UpdateVocabularyWithCategories(*Vocabulary, []string) (*Vocabulary, error)
	DeleteVocabularyById(uint) error
}
