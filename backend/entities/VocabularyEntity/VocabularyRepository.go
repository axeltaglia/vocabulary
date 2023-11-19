package VocabularyEntity

type VocabularyRepository interface {
	CreateVocabulary(vocabulary *Vocabulary) (*Vocabulary, error)
	CreateVocabularyWithCategories(*Vocabulary, []string) (*Vocabulary, error)
	GetAllVocabulariesWithCategories() ([]Vocabulary, error)
	FindVocabularyById(id uint) (*Vocabulary, error)
	FindCategories() []Category
	FindCategoriesByVocabularyId(id uint) ([]Category, error)
	UpdateVocabulary(*Vocabulary) (*Vocabulary, error)
	DisassociateCategoriesFromVocabulary(*Vocabulary) error
	AssociateCategoryToVocabulary(*Vocabulary, *Category) (*Vocabulary, error)
	CreateCategoryIfNotExist(categoryName string) (*Category, error)
	DeleteVocabularyById(uint) error
}
