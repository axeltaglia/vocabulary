package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"vocabulary/modules"
	"vocabulary/util"
)

type Endpoints struct {
	router *gin.Engine
	db     *gorm.DB
}

func (o *Endpoints) createVocabulary(c *gin.Context, tx *gorm.DB) {
	var vocabulary modules.Vocabulary
	err := c.BindJSON(&vocabulary)
	util.CheckErr(err)
	tx.Create(&vocabulary)
	tx.Find(&vocabulary)
	c.JSON(http.StatusCreated, vocabulary)
}

func (o *Endpoints) getVocabularies(c *gin.Context, tx *gorm.DB) {
	var vocabularies []modules.Vocabulary
	tx.Order("created_at DESC").Find(&vocabularies)

	c.JSON(http.StatusOK, vocabularies)
}

func (o *Endpoints) getVocabulary(c *gin.Context, tx *gorm.DB) {
	id := c.Params.ByName("id")
	var vocabulary modules.Vocabulary
	tx.First(&vocabulary, id)

	if vocabulary.Id != 0 {
		c.JSON(http.StatusOK, vocabulary)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (o *Endpoints) getVocabularyCategories(c *gin.Context, tx *gorm.DB) {
	id := c.Params.ByName("id")
	var vocabulary modules.Vocabulary
	tx.Model(&vocabulary).First(&vocabulary, id).Association("Categories").Find(&vocabulary.Categories)

	if vocabulary.Id != 0 {
		c.JSON(http.StatusOK, vocabulary.Categories)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (o *Endpoints) updateVocabulary(c *gin.Context, tx *gorm.DB) {
	id := c.Params.ByName("id")
	var vocabulary modules.Vocabulary
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
	vocabulary := &modules.Vocabulary{}
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
			category := modules.Category{}
			if err := tx.Where("name = ?", categoryName).First(&category).Error; err != nil {
				category = modules.Category{Name: categoryName}
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
	var vocabulary modules.Vocabulary
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
		category := modules.Category{}
		if err := tx.Where("name = ?", categoryName).First(&category).Error; err != nil {
			category = modules.Category{Name: categoryName}
			tx.Create(&category)
		}
		tx.Model(&vocabulary).Association("Categories").Append(category)
	}

	c.JSON(http.StatusCreated, vocabulary)
}

func (o *Endpoints) getCategories(c *gin.Context, tx *gorm.DB) {
	var categories []modules.Category
	tx.Order("name").Find(&categories)

	c.JSON(http.StatusOK, categories)
}

func (o *Endpoints) Handle() {
	/*
		o.router.GET("/ping", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "PONG",
			})
		})

	*/
	o.handleWithTx("/getVocabularies", o.getVocabularies)
	o.handleWithTx("/getVocabulary/:id", o.getVocabulary)
	o.handleWithTx("/getVocabularyCategories/:id", o.getVocabularyCategories)
	o.handleWithTx("/createVocabulary", o.createVocabulary)
	o.handleWithTx("/updateVocabulary/:id", o.updateVocabulary)
	o.handleWithTx("/updateVocabularyWithCategories/:id", o.updateVocabularyWithCategories)
	o.handleWithTx("/deleteVocabulary/:id", o.deleteVocabulary)
	o.handleWithTx("/createVocabularyWithCategories", o.createVocabularyWithCategories)
	o.handleWithTx("/getCategories", o.getCategories)
}

func (o *Endpoints) handleWithTx(relativePath string, f func(c *gin.Context, tx *gorm.DB)) {
	o.router.POST(relativePath, func(c *gin.Context) {
		tx := o.db.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r)
			}
		}()

		f(c, tx)

		if tx.Error != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database transaction failed"})
		} else {
			tx.Commit()
		}
	})
}

func NewEndpoints(router *gin.Engine, db *gorm.DB) *Endpoints {
	return &Endpoints{
		router: router,
		db:     db,
	}
}
