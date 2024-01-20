package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) updateVocabularyWithCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	var requestData VocabularyWithCategoriesRequest
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vocabulary, err := vocabularyEntity.UpdateWithCategories(requestData.Vocabulary.MapToEntity(), requestData.CategoryNames)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var response UpdateVocabularyWithCategoriesResponse
	response.MapFromEntity(vocabulary)
	c.JSON(http.StatusOK, response)
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
