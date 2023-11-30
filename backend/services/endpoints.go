package services

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"vocabulary/entities"
	"vocabulary/entities/VocabularyEntity"
)

type Endpoints struct {
	router            *gin.Engine
	apiPort           string
	repositoryFactory entities.RepositoryFactory
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

	vocabulary := vocabularyEntity.Create(request.MapToEntity())

	var response Vocabulary
	response.MapFromEntity(vocabulary)

	c.JSON(http.StatusCreated, response)
}

func (o *Endpoints) getVocabularies(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	vocabularies := vocabularyEntity.GetAllVocabulariesWithCategories()

	var getVocabulariesResponse GetVocabulariesResponse
	getVocabulariesResponse.MapFromEntities(vocabularies)

	c.JSON(http.StatusOK, getVocabulariesResponse.Vocabularies)
}

func (o *Endpoints) getVocabulary(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	strId := c.Params.ByName("id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		panic("Id must be a number")
	}

	vocabulary := vocabularyEntity.GetVocabulary(uint(id))

	var response Vocabulary
	response.MapFromEntity(vocabulary)

	c.JSON(http.StatusOK, response)
}

func (o *Endpoints) getVocabularyCategories(c *gin.Context, vocabularyEntity VocabularyEntity.Entity) {
	strId := c.Params.ByName("id")
	id, err := strconv.ParseUint(strId, 10, 32)
	if err != nil {
		return
	}

	entityCategories := vocabularyEntity.GetCategoriesFromVocabulary(uint(id))

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

/*

func (o *Endpoints) getVocabularyCategories(c *gin.Context, tx *gorm.DB) {
	id := c.Params.ByName("id")
	var vocabulary Entity.Vocabulary
	tx.Model(&vocabulary).First(&vocabulary, id).Association("Categories").Find(&vocabulary.Categories)

	if vocabulary.Id != 0 {
		c.JSON(http.StatusOK, vocabulary.Categories)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (o *Endpoints) updateVocabulary(c *gin.Context, tx *gorm.DB) {
	id := c.Params.ByName("id")
	var vocabulary Entity.Vocabulary
	tx.First(&vocabulary, id)

	if vocabulary.Id != 0 {
		err := c.BindJSON(&vocabulary)
		util.CheckErr(err)
		tx.Save(&vocabulary)
		c.JSON(http.StatusOK, vocabulary)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (o *Endpoints) updateVocabularyWithCategories(c *gin.Context, tx *gorm.DB) {
	id := c.Params.ByName("id")
	vocabulary := &Entity.Vocabulary{}
	tx.First(vocabulary, id)

	if vocabulary.Id != 0 {
		var requestData VocabularyWithCategories
		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		vocabulary = &requestData.Vocabulary
		tx.Save(vocabulary)

		tx.Model(vocabulary).Association("Categories").Clear()

		for _, categoryName := range requestData.Categories {
			category := Entity.Category{}
			if err := tx.Where("name = ?", categoryName).First(&category).Error; err != nil {
				category = Entity.Category{Name: categoryName}
				tx.Create(&category)
			}
			tx.Model(vocabulary).Association("Categories").Append(category)
		}

		c.JSON(http.StatusOK, vocabulary)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

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

func (o *Endpoints) getCategories(c *gin.Context, tx *gorm.DB) {
	var categories []Entity.Category
	tx.Order("created_at DESC").Find(&categories)

	c.JSON(http.StatusOK, categories)
}

*/

func (o *Endpoints) handle() {
	o.handleTxWithVocabularyEntity("/getVocabularies", o.getVocabularies)
	o.handleTxWithVocabularyEntity("/getVocabulary/:id", o.getVocabulary)
	o.handleTxWithVocabularyEntity("/getVocabularyCategories/:id", o.getVocabularyCategories)
	o.handleTxWithVocabularyEntity("/createVocabulary", o.createVocabulary)
	//o.handleWithTx("/updateVocabulary/:id", o.updateVocabulary)
	//o.handleWithTx("/updateVocabularyWithCategories/:id", o.updateVocabularyWithCategories)
	//o.handleWithTx("/deleteVocabulary/:id", o.deleteVocabulary)
	//o.handleWithTx("/createVocabularyWithCategories", o.createVocabularyWithCategories)
	o.handleTxWithVocabularyEntity("/getCategories", o.getCategories)
}

func (o *Endpoints) handleTxWithVocabularyEntity(relativePath string, f func(c *gin.Context, vocabularyEntity VocabularyEntity.Entity)) {
	o.router.POST(relativePath, func(c *gin.Context) {
		o.repositoryFactory.BeginTransaction()
		defer func() {
			if r := recover(); r != nil {
				o.repositoryFactory.RollbackTransaction()
				panic(r)
			}
		}()

		vocabularyRepository := o.repositoryFactory.GetVocabularyRepository()
		vocabularyEntity := VocabularyEntity.New(vocabularyRepository)
		f(c, vocabularyEntity)

		//if tx.Error != nil {
		//tx.Rollback()
		//c.JSON(http.StatusInternalServerError, gin.H{"error": "Database transaction failed"})
		//} else {
		o.repositoryFactory.CommitTransaction()
		//}
	})
}

func (o *Endpoints) ListenAndServe() {
	o.handle()
	err := http.ListenAndServe(":"+o.apiPort, o.router)
	if err != nil {
		panic(err)
	}
}

func NewEndpoints(apiPort string, repositoryFactory entities.RepositoryFactory) *Endpoints {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(corsConfig))
	return &Endpoints{
		apiPort:           apiPort,
		router:            router,
		repositoryFactory: repositoryFactory,
	}
}
