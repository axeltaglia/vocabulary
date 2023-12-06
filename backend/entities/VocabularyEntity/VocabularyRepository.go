package VocabularyEntity

type VocabularyRepository interface {
	CreateVocabulary(vocabulary Vocabulary) Vocabulary
	GetAllVocabulariesWithCategories() []Vocabulary
	FindVocabularyById(id uint) Vocabulary
	FindCategories() []Category
	UpdateVocabulary(vocabulary Vocabulary) Vocabulary
	UpdateVocabularyWithCategories(Vocabulary, []string)
}
