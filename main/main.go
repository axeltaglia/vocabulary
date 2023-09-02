package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"vocabulary/modules"
	"vocabulary/services"
)

// docker run --name pg_vocabulary_ctn -e POSTGRES_USER=vocabulary -e POSTGRES_PASSWORD=vocabulary -e POSTGRES_DB=vocabulary -p 5435:5432 -d postgres
func main() {
	// Initialize the database
	var db *gorm.DB
	initDB(db)

	// Initialize the Gin router
	router := gin.Default()

	// Define API routes
	endpoints := services.NewEndpoints(router, db)
	endpoints.Handle()

	// Start the server
	router.Run(":8080")
}

// Initialize the database
func initDB(db *gorm.DB) {
	var err error
	db, err = gorm.Open("postgres", "host=localhost port=5435 user=vocabulary dbname=vocabulary password=vocabulary sslmode=disable")
	if err != nil {
		panic("Failed to connect to database")
	}

	db.AutoMigrate(&modules.Vocabulary{})
}
