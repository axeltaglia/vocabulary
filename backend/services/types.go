package services

import (
	"vocabulary/entities/VocabularyEntity"
)

type CreateVocabularyRequest struct {
	Words        *string `json:"words"`
	Translation  *string `json:"translation"`
	UsedInPhrase *string `json:"usedInPhrase"`
	Explanation  *string `json:"explanation"`
}

func (o *CreateVocabularyRequest) IsValid() bool {
	if o.Words == nil || *o.Words == "" {
		return false
	}

	return true
}

func (o *CreateVocabularyRequest) MapToEntity() VocabularyEntity.Vocabulary {
	return VocabularyEntity.Vocabulary{
		Words:        o.Words,
		Translation:  o.Translation,
		UsedInPhrase: o.UsedInPhrase,
		Explanation:  o.Explanation,
	}
}

type Vocabulary struct {
	Id           *uint      `json:"id"`
	Words        *string    `json:"words"`
	Translation  *string    `json:"translation"`
	UsedInPhrase *string    `json:"usedInPhrase"`
	Explanation  *string    `json:"explanation"`
	Categories   []Category `json:"categories"`
}

func (o *Vocabulary) MapFromEntity(vocabulary VocabularyEntity.Vocabulary) {
	o.Id = vocabulary.Id
	o.Words = vocabulary.Words
	o.Translation = vocabulary.Translation
	o.UsedInPhrase = vocabulary.UsedInPhrase
	o.Explanation = vocabulary.Explanation
	o.Categories = MapCategoriesFromEntity(vocabulary.Categories)
}
func MapCategoriesFromEntity(entityCategories []VocabularyEntity.Category) []Category {
	var categories []Category
	for _, entityCategory := range entityCategories {
		categories = append(categories, Category{
			Id:   entityCategory.Id,
			Name: entityCategory.Name,
		})
	}
	return categories
}

type Category struct {
	Id   *uint   `json:"id"`
	Name *string `json:"name"`
}

type GetVocabulariesResponse struct {
	Vocabularies []Vocabulary `json:"vocabularies"`
}

func (o *GetVocabulariesResponse) MapFromEntities(vocabularies []VocabularyEntity.Vocabulary) {
	for _, entityVocabulary := range vocabularies {
		var categories []Category
		for _, entityCategory := range entityVocabulary.Categories {
			categories = append(categories, Category{
				Id:   entityCategory.Id,
				Name: entityCategory.Name,
			})
		}

		o.Vocabularies = append(o.Vocabularies, Vocabulary{
			Id:           entityVocabulary.Id,
			Words:        entityVocabulary.Words,
			Translation:  entityVocabulary.Translation,
			UsedInPhrase: entityVocabulary.UsedInPhrase,
			Explanation:  entityVocabulary.Explanation,
			Categories:   categories,
		})
	}
}

type VocabularyWithCategories struct {
	Vocabulary VocabularyEntity.Vocabulary `json:"vocabulary"`
	Categories []string                    `json:"categories"`
}

type GetVocabularyCategoriesResponse struct {
	Categories []Category `json:"categories"`
}

func (o *GetVocabularyCategoriesResponse) MapFromEntities(entityCategories []VocabularyEntity.Category) {
	var categories []Category
	for _, entityCategory := range entityCategories {
		categories = append(categories, Category{
			Id:   entityCategory.Id,
			Name: entityCategory.Name,
		})
	}

	o.Categories = categories
}
