package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) updateVocabulary(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	var request UpdateVocabularyRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	vocabulary, err := vocabularyEntity.Update(request.MapToEntity())
	var response Vocabulary
	response.MapFromEntity(vocabulary)
	c.JSON(http.StatusOK, response)
}

type UpdateVocabularyRequest struct {
	Words        *string `json:"words"`
	Translation  *string `json:"translation"`
	UsedInPhrase *string `json:"usedInPhrase"`
	Explanation  *string `json:"explanation"`
}

func (o *UpdateVocabularyRequest) IsValid() bool {
	if o.Words == nil || *o.Words == "" {
		return false
	}

	return true
}

func (o *UpdateVocabularyRequest) MapToEntity() *VocabularyEntity.Vocabulary {
	return &VocabularyEntity.Vocabulary{
		Words:        o.Words,
		Translation:  o.Translation,
		UsedInPhrase: o.UsedInPhrase,
		Explanation:  o.Explanation,
	}
}
