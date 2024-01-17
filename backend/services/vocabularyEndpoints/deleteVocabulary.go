package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) deleteVocabulary(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	strId := c.Params.ByName("id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if err = vocabularyEntity.Delete(uint(id)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}
}
