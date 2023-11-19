package vocabularyEndpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vocabulary/entities/VocabularyEntity"
)

func (o *Endpoints) deleteVocabulary(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) error {
	id, err := getIdFromRequest(c)
	if err != nil {
		return err
	}

	if err = vocabularyEntity.Delete(uint(id)); err != nil {
		return APIError{
			Msg:         "The Vocabulary couldn't be deleted",
			Status:      http.StatusBadRequest,
			originalErr: err,
		}
	}

	return nil
}
