package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) updateVocabularyWithCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) error {
	var requestData VocabularyWithCategoriesRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		return APIError{
			Msg:         "Error",
			Status:      http.StatusBadRequest,
			originalErr: err,
		}
	}

	vocabulary, err := vocabularyEntity.UpdateWithCategories(requestData.Vocabulary.MapToEntity(), requestData.CategoryNames)
	if err != nil {
		return APIError{
			Msg:         "Error",
			Status:      http.StatusBadRequest,
			originalErr: err,
		}
	}

	var response UpdateVocabularyWithCategoriesResponse
	response.MapFromEntity(vocabulary)
	c.JSON(http.StatusOK, response)
	return nil
}

type VocabularyWithCategoriesRequest struct {
	Vocabulary    Vocabulary `json:"vocabulary"`
	CategoryNames []string   `json:"categoryNames"`
}

type UpdateVocabularyWithCategoriesResponse struct {
	Vocabulary
}

func (o *UpdateVocabularyWithCategoriesResponse) MapFromEntity(vocabulary *VocabularyEntity.Vocabulary) {
	o.Vocabulary.MapFromEntity(vocabulary)
}
