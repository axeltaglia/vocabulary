package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/logger"
)

func (o *Endpoints) createVocabulary(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	var request CreateVocabularyRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if !request.IsValid() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Words is mandatory"})
		return
	}

	vocabulary, err := vocabularyEntity.Create(request.MapToEntity())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.LogInfo("vocabulary/createVocabulary has failed")
		return
	}

	var createVocabularyResponse CreateVocabularyResponse
	createVocabularyResponse.MapFromEntity(vocabulary)
	c.JSON(http.StatusCreated, createVocabularyResponse)
}

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

func (o *CreateVocabularyRequest) MapToEntity() *VocabularyEntity.Vocabulary {
	return &VocabularyEntity.Vocabulary{
		Words:        o.Words,
		Translation:  o.Translation,
		UsedInPhrase: o.UsedInPhrase,
		Explanation:  o.Explanation,
	}
}

type CreateVocabularyResponse struct {
	Vocabulary
}

func (o *CreateVocabularyResponse) MapFromEntity(vocabulary *VocabularyEntity.Vocabulary) {
	o.Vocabulary.MapFromEntity(vocabulary)
}
