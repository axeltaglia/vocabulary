package vocabularyEndpoints

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/logger"
)

func (o *Endpoints) createVocabularyWithCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	var request CreateVocabularyWithCategoriesRequest
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := request.Validate(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vocabulary, err := vocabularyEntity.CreateWithCategories(request.MapToEntity(), request.CategoryNames)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		logger.LogInfo("vocabulary/createVocabulary has failed")
		return
	}

	var createVocabularyResponse CreateVocabularyWithCategoriesResponse
	createVocabularyResponse.MapFromEntity(vocabulary)
	c.JSON(http.StatusCreated, createVocabularyResponse)
}

type CreateVocabularyWithCategoriesRequest struct {
	Vocabulary    Vocabulary `json:"vocabulary"`
	CategoryNames []string   `json:"categoryNames"`
}

func (o *CreateVocabularyWithCategoriesRequest) Validate() error {
	if o.Vocabulary.Words == nil || *o.Vocabulary.Words == "" {
		return errors.New("field words mandatory")
	}

	return nil
}

func (o *CreateVocabularyWithCategoriesRequest) MapToEntity() *VocabularyEntity.Vocabulary {
	return &VocabularyEntity.Vocabulary{
		Words:        o.Vocabulary.Words,
		Translation:  o.Vocabulary.Translation,
		UsedInPhrase: o.Vocabulary.UsedInPhrase,
		Explanation:  o.Vocabulary.Explanation,
	}
}

type CreateVocabularyWithCategoriesResponse struct {
	Vocabulary
}

func (o *CreateVocabularyWithCategoriesResponse) MapFromEntity(vocabulary *VocabularyEntity.Vocabulary) {
	o.Vocabulary.MapFromEntity(vocabulary)
}
