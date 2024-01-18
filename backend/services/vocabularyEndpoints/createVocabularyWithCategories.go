package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/logger"
)

func (o *Endpoints) createVocabularyWithCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	var request CreateVocabularyWithCategoriesRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if !request.IsValid() {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Words is mandatory"})
		return
	}

	vocabulary, err := vocabularyEntity.CreateWithCategories(request.MapToEntity(), request.Categories)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.LogInfo("vocabulary/createVocabulary has failed")
		return
	}

	var createVocabularyResponse CreateVocabularyResponse
	createVocabularyResponse.MapFromEntity(vocabulary)
	c.JSON(http.StatusCreated, createVocabularyResponse)
}

type CreateVocabularyWithCategoriesRequest struct {
	Vocabulary Vocabulary `json:"vocabulary"`
	Categories []string   `json:"categories"`
}

func (o *CreateVocabularyWithCategoriesRequest) IsValid() bool {
	if o.Vocabulary.Words == nil || *o.Vocabulary.Words == "" {
		return false
	}

	return true
}

func (o CreateVocabularyWithCategoriesRequest) MapToEntity() *VocabularyEntity.Vocabulary {
	return &VocabularyEntity.Vocabulary{
		Words:        o.Vocabulary.Words,
		Translation:  o.Vocabulary.Translation,
		UsedInPhrase: o.Vocabulary.UsedInPhrase,
		Explanation:  o.Vocabulary.Explanation,
	}
}
