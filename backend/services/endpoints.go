package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"
	"vocabulary/entities"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/logger"
)

type Endpoints struct {
	router              *gin.Engine
	txRepositoryHandler entities.TxRepositoryHandler
}

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

	var response Vocabulary
	response.MapFromEntity(vocabulary)
	c.JSON(http.StatusCreated, response)
}

func (o *Endpoints) updateVocabulary(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	var request UpdateVocabularyRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	vocabulary, err := vocabularyEntity.Update(request.MapToEntity())
	var response Vocabulary
	response.MapFromEntity(vocabulary)
	c.JSON(http.StatusOK, response)
}

func (o *Endpoints) getVocabularies(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	vocabularies, err := vocabularyEntity.GetAllVocabulariesWithCategories()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
	}

	var getVocabulariesResponse GetVocabulariesResponse
	getVocabulariesResponse.MapFromEntities(vocabularies)

	c.JSON(http.StatusOK, getVocabulariesResponse.Vocabularies)
}

func (o *Endpoints) getVocabulary(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	strId := c.Params.ByName("id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	vocabulary, err := vocabularyEntity.GetVocabulary(uint(id))

	var response Vocabulary
	response.MapFromEntity(vocabulary)

	c.JSON(http.StatusOK, response)
}

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

func (o *Endpoints) getCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	entityCategories := vocabularyEntity.GetAllCategories()

	var getCategoriesResponse GetCategoriesResponse
	getCategoriesResponse.MapFromEntities(entityCategories)

	c.JSON(http.StatusOK, getCategoriesResponse.Categories)
}

func (o *Endpoints) updateVocabularyWithCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	strId := c.Params.ByName("id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	vocabulary, err := vocabularyEntity.GetVocabulary(uint(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request format"})
		return
	}

	if vocabulary.Id != nil {
		var requestData VocabularyWithCategories
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		vocabulary, err := vocabularyEntity.UpdateWithCategories(requestData.MapToEntity(), requestData.Categories)
		if err != nil {
			return
		}

		c.JSON(http.StatusOK, vocabulary)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

/*
func (o *Endpoints) deleteVocabulary(c *gin.Context, tx *gorm.DB) {
	id := c.Params.ByName("id")
	var vocabulary Entity.Vocabulary
	tx.First(&vocabulary, id)

	if vocabulary.Id != 0 {
		tx.Delete(&vocabulary)
		c.JSON(http.StatusOK, gin.H{"message": "Vocabulary deleted"})
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (o *Endpoints) createVocabularyWithCategories(c *gin.Context, tx *gorm.DB) {
	var requestData VocabularyWithCategories
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	vocabulary := requestData.Vocabulary
	tx.Create(&vocabulary)

	for _, categoryName := range requestData.Categories {
		category := Entity.Category{}
		if err := tx.Where("name = ?", categoryName).First(&category).Error; err != nil {
			category = Entity.Category{Name: categoryName}
			tx.Create(&category)
		}
		tx.Model(&vocabulary).Association("Categories").Append(category)
	}

	c.JSON(http.StatusCreated, vocabulary)
}

*/

func (o *Endpoints) handle() {
	o.handleTxWithVocabularyEntity("/getVocabularies", o.getVocabularies)
	o.handleTxWithVocabularyEntity("/getVocabulary/:id", o.getVocabulary)
	o.handleTxWithVocabularyEntity("/getVocabularyCategories/:id", o.getVocabularyCategories)
	o.handleTxWithVocabularyEntity("/createVocabulary", o.createVocabulary)
	o.handleTxWithVocabularyEntity("/updateVocabulary/:id", o.updateVocabulary)
	o.handleTxWithVocabularyEntity("/updateVocabularyWithCategories/:id", o.updateVocabularyWithCategories)
	//o.handleWithTx("/deleteVocabulary/:id", o.deleteVocabulary)
	//o.handleWithTx("/createVocabularyWithCategories", o.createVocabularyWithCategories)
	o.handleTxWithVocabularyEntity("/getCategories", o.getCategories)
}

func (o *Endpoints) handleTxWithVocabularyEntity(relativePath string, f func(c *gin.Context, vocabularyEntity VocabularyEntity.Entity)) {
	o.router.POST(relativePath, func(c *gin.Context) {
		txRepositoryFactory := o.txRepositoryHandler.GetTxRepositoryFactory()

		defer func() {
			if r := recover(); r != nil {
				txRepositoryFactory.RollbackTransaction()
				stackTrace := string(debug.Stack())
				fmt.Printf("%v\n%s\n", r, stackTrace)
			}
		}()

		vocabularyRepository := txRepositoryFactory.GetVocabularyRepository()
		vocabularyEntity := VocabularyEntity.New(vocabularyRepository)
		f(c, vocabularyEntity)

		if txRepositoryFactory.TransactionError() != nil {
			txRepositoryFactory.RollbackTransaction()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database transaction failed"})
		} else {
			txRepositoryFactory.CommitTransaction()
		}
	})
}

func Marshal(object interface{}) []byte {
	bytes, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}
	return bytes
}

func (o *Endpoints) ListenAndServe(apiPort string) {
	o.handle()
	err := http.ListenAndServe(":"+apiPort, o.router)
	if err != nil {
		panic(err)
	}
}

func NewEndpoints(txRepositoryHandler entities.TxRepositoryHandler) *Endpoints {
	router := gin.Default()
	router.Use(corsMiddleware())
	router.Use(loggerMiddleware())
	return &Endpoints{
		router:              router,
		txRepositoryHandler: txRepositoryHandler,
	}
}

func corsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	return cors.New(corsConfig)
}

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		logger.Log().WithFields(logrus.Fields{
			"status":    statusCode,
			"latency":   latency,
			"client_ip": clientIP,
			"method":    method,
			"path":      path,
			"raw_query": raw,
			"error":     errorMessage,
		}).Info("Request handled")
	}
}
