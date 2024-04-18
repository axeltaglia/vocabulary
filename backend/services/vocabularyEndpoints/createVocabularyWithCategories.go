package vocabularyEndpoints

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"net/http"
	"strings"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) createVocabularyWithCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) error {
	var request CreateVocabularyWithCategoriesRequest
	if err := c.BindJSON(&request); err != nil {
		return APIError{
			Msg:         "Invalid request format",
			Status:      http.StatusBadRequest,
			originalErr: err,
		}
	}

	if err := request.Validate(); err != nil {
		return APIError{
			Msg:         "Invalid request format",
			Status:      http.StatusBadRequest,
			originalErr: err,
		}
	}

	vocabulary, err := vocabularyEntity.CreateWithCategories(request.MapToEntity(), request.CategoryNames)
	if err != nil {
		return APIError{
			Msg:         "Invalid request format",
			Status:      http.StatusBadRequest,
			originalErr: err,
		}
	}

	var createVocabularyResponse CreateVocabularyWithCategoriesResponse
	createVocabularyResponse.MapFromEntity(vocabulary)
	c.JSON(http.StatusCreated, createVocabularyResponse)
	return nil
}

type CreateVocabularyWithCategoriesRequest struct {
	Vocabulary    Vocabulary `json:"vocabulary" validate:"required"`
	CategoryNames []string   `json:"categoryNames" validate:"dive"`
}

func (o *CreateVocabularyWithCategoriesRequest) Validate() error {
	v := validator.New()

	if err := v.Struct(o); err != nil {
		errMsg := ""
		for _, e := range err.(validator.ValidationErrors) {
			errMsg += fmt.Sprintf("Field %s failed validation with tag %s. Custom message: %s\n", e.Field(), e.Tag(), e.Param()) + ". "
		}
		return errors.New(errMsg)
	}

	if err := o.ValidateCategories(); err != nil {
		return err
	}

	return nil
}

func (o *CreateVocabularyWithCategoriesRequest) ValidateCategories() error {
	var categoryNamesTrimed []string
	for _, category := range o.CategoryNames {
		if len(category) < 30 {
			errMessage := "there is at least one category with more than 30 characters length"
			return errors.New(errMessage)
		}
		categoryNamesTrimed = append(categoryNamesTrimed, strings.TrimSpace(category))
	}
	o.CategoryNames = categoryNamesTrimed
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
