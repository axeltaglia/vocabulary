package services

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"vocabulary/modules"
)

type Endpoints struct {
	router *gin.Engine
	db     *gorm.DB
}

// Create a new Vocabulary
func (o *Endpoints) createVocabulary(c *gin.Context) {
	var vocabulary modules.Vocabulary
	c.BindJSON(&vocabulary)

	o.db.Create(&vocabulary)
	o.db.Find(&vocabulary)
	c.JSON(http.StatusCreated, vocabulary)
}

// Get all Vocabularies
func (o *Endpoints) getVocabularies(c *gin.Context) {
	var vocabularies []modules.Vocabulary
	o.db.Find(&vocabularies)

	c.JSON(http.StatusOK, vocabularies)
}

// Get a single Vocabulary by ID
func (o *Endpoints) getVocabulary(c *gin.Context) {
	id := c.Params.ByName("id")
	var vocabulary modules.Vocabulary
	o.db.First(&vocabulary, id)

	if vocabulary.ID != 0 {
		c.JSON(http.StatusOK, vocabulary)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Update a Vocabulary
func (o *Endpoints) updateVocabulary(c *gin.Context) {
	id := c.Params.ByName("id")
	var vocabulary modules.Vocabulary
	o.db.First(&vocabulary, id)

	if vocabulary.ID != 0 {
		c.BindJSON(&vocabulary)
		o.db.Save(&vocabulary)
		c.JSON(http.StatusOK, vocabulary)
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

// Delete a Vocabulary
func (o *Endpoints) deleteVocabulary(c *gin.Context) {
	id := c.Params.ByName("id")
	var vocabulary modules.Vocabulary
	o.db.First(&vocabulary, id)

	if vocabulary.ID != 0 {
		o.db.Delete(&vocabulary)
		c.JSON(http.StatusOK, gin.H{"message": "Vocabulary deleted"})
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func (o *Endpoints) Handle() {
	o.router.GET("/vocabularies", o.getVocabularies)
	o.router.GET("/vocabularies/:id", o.getVocabulary)
	o.router.POST("/vocabularies", o.createVocabulary)
	o.router.PUT("/vocabularies/:id", o.updateVocabulary)
	o.router.DELETE("/vocabularies/:id", o.deleteVocabulary)
}

func NewEndpoints(router *gin.Engine, db *gorm.DB) *Endpoints {
	return &Endpoints{
		router: router,
		db:     db,
	}
}
