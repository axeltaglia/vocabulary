package gormRepository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"os"
	"time"
	"vocabulary/gormRepository/VocabularyGormRepository"
)

type DbConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func InitDb(dbConfig DbConfig) *gorm.DB {
	maxAttempts := 30
	attempt := 1
	for {
		var db *gorm.DB
		var err error
		args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Database, dbConfig.Password)
		db, err = gorm.Open("postgres", args)
		log.Println(args)
		if err == nil {
			// AutoMigrate both Vocabulary and VocabularyCategory models
			db.AutoMigrate(&VocabularyGormRepository.Vocabulary{})
			db.AutoMigrate(&VocabularyGormRepository.Category{})

			// Create a migration to define the join table with foreign keys
			db.Table("vocabulary_categories").AddForeignKey("vocabulary_id", "vocabularies(id)", "CASCADE", "CASCADE")
			db.Table("vocabulary_categories").AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")

			return db
		}
		attempt++

		if attempt > maxAttempts {
			fmt.Println("Max connection attempts reached. Exiting.")
			os.Exit(1)
		}
		fmt.Println("Retrying in 1 second...")
		time.Sleep(1 * time.Second)
	}
}
