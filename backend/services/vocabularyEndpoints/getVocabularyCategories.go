package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) getVocabularyCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	strId := c.Params.ByName("id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	entityCategories, err := vocabularyEntity.GetCategoriesFromVocabulary(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	var getVocabularyCategoriesResponse GetVocabularyCategoriesResponse
	getVocabularyCategoriesResponse.MapFromEntities(entityCategories)

	c.JSON(http.StatusOK, getVocabularyCategoriesResponse.Categories)
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
