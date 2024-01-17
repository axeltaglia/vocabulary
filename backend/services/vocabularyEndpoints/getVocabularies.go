package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) getVocabularies(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	vocabularies, err := vocabularyEntity.GetAllVocabulariesWithCategories()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
	}

	var getVocabulariesResponse GetVocabulariesResponse
	getVocabulariesResponse.MapFromEntities(vocabularies)

	c.JSON(http.StatusOK, getVocabulariesResponse.Vocabularies)
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
