package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) getVocabularies(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) error {
	vocabularies, err := vocabularyEntity.GetAllVocabulariesWithCategories()
	if err != nil {
		return APIError{
			Msg:         "Invalid request format",
			Status:      http.StatusBadRequest,
			originalErr: err,
		}
	}

	var getVocabulariesResponse GetVocabulariesResponse
	getVocabulariesResponse.MapFromEntities(vocabularies)

	c.JSON(http.StatusOK, getVocabulariesResponse.Vocabularies)
	return nil
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
