package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) getCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) error {
	entityCategories := vocabularyEntity.GetAllCategories()

	var getCategoriesResponse GetCategoriesResponse
	getCategoriesResponse.MapFromEntities(entityCategories)

	c.JSON(http.StatusOK, getCategoriesResponse.Categories)
	return nil
}

type GetCategoriesResponse struct {
	Categories []Category `json:"categories"`
}

func (o *GetCategoriesResponse) MapFromEntities(entityCategories []VocabularyEntity.Category) {
	var categories []Category
	for _, entityCategory := range entityCategories {
		categories = append(categories, Category{
			Id:   entityCategory.Id,
			Name: entityCategory.Name,
		})
	}

	o.Categories = categories
}
