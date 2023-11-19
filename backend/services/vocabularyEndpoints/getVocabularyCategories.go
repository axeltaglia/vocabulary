package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) getVocabularyCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) error {
	strId := c.Params.ByName("id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		return APIError{
			Msg:         "Invalid request format",
			Status:      http.StatusBadRequest,
			originalErr: err,
		}
	}

	entityCategories, err := vocabularyEntity.GetCategoriesFromVocabulary(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return APIError{
			Msg:         "Error",
			Status:      http.StatusBadRequest,
			originalErr: err,
		}
	}

	var getVocabularyCategoriesResponse GetVocabularyCategoriesResponse
	getVocabularyCategoriesResponse.MapFromEntities(entityCategories)

	c.JSON(http.StatusOK, getVocabularyCategoriesResponse.Categories)

	return nil
}

type GetVocabularyCategoriesResponse struct {
	Categories []Category `json:"categories"`
}

func (o *GetVocabularyCategoriesResponse) MapFromEntities(entityCategories []VocabularyEntity.Category) {
	var categories []Category
	for _, entityCategory := range entityCategories {
		categories = append(categories, Category{
			Id:   entityCategory.Id,
			Name: entityCategory.Name,
		})
	}
	o.Categories = categories
}
