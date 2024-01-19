package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) updateVocabularyWithCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	strId := c.Params.ByName("id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	vocabulary, err := vocabularyEntity.GetVocabulary(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if vocabulary.Id != nil {
		var requestData VocabularyWithCategoriesRequest
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		vocabulary, err := vocabularyEntity.UpdateWithCategories(requestData.MapToEntity(), requestData.Categories)
		if err != nil {
			return
		}

		var response UpdateVocabularyWithCategoriesResponse
		response.MapFromEntity(vocabulary)

		c.JSON(http.StatusOK, response)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

type VocabularyWithCategoriesRequest struct {
	Vocabulary VocabularyEntity.Vocabulary `json:"vocabulary"`
	Categories []string                    `json:"categories"`
}

func (o VocabularyWithCategoriesRequest) MapToEntity() *VocabularyEntity.Vocabulary {
	return &VocabularyEntity.Vocabulary{
		Id:           o.Vocabulary.Id,
		Words:        o.Vocabulary.Words,
		Translation:  o.Vocabulary.Translation,
		UsedInPhrase: o.Vocabulary.UsedInPhrase,
		Explanation:  o.Vocabulary.Explanation,
	}
}

type UpdateVocabularyWithCategoriesResponse struct {
	Vocabulary
}

func (o *UpdateVocabularyWithCategoriesResponse) MapFromEntity(vocabulary *VocabularyEntity.Vocabulary) {
	o.Vocabulary.MapFromEntity(vocabulary)
}
