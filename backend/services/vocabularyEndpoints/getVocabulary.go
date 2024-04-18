package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) getVocabulary(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) error {
	id, err := getIdFromRequest(c)
	if err != nil {
		return err
	}

	vocabulary, err := vocabularyEntity.GetVocabulary(uint(id))

	var response GetVocabularyResponse
	response.MapFromEntity(vocabulary)

	c.JSON(http.StatusOK, response)
	return nil
}

type GetVocabularyResponse struct {
	Id           *uint      `json:"id"`
	Words        *string    `json:"words"`
	Translation  *string    `json:"translation"`
	UsedInPhrase *string    `json:"usedInPhrase"`
	Explanation  *string    `json:"explanation"`
	Categories   []Category `json:"categories"`
}

func (o *GetVocabularyResponse) MapFromEntity(vocabulary *VocabularyEntity.Vocabulary) {
	o.Id = vocabulary.Id
	o.Words = vocabulary.Words
	o.Translation = vocabulary.Translation
	o.UsedInPhrase = vocabulary.UsedInPhrase
	o.Explanation = vocabulary.Explanation
	o.Categories = MapCategoriesFromEntity(vocabulary.Categories)
}

func (o *GetVocabularyResponse) MapToEntity() VocabularyEntity.Vocabulary {
	return VocabularyEntity.Vocabulary{
		Id:           o.Id,
		Words:        o.Words,
		Translation:  o.Translation,
		UsedInPhrase: o.UsedInPhrase,
		Explanation:  o.Explanation,
	}
}

func MapCategoriesFromEntity(entityCategories []VocabularyEntity.Category) []Category {
	var categories []Category
	for _, entityCategory := range entityCategories {
		categories = append(categories, Category{
			Id:   entityCategory.Id,
			Name: entityCategory.Name,
		})
	}
	return categories
}
