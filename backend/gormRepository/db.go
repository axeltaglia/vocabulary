package gormRepository

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
	"vocabulary/gormRepository/VocabularyGormRepository"
)

type DbConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DbName   string `json:"dbName"`
}

func ConnectToDbWithMaxAttempts(dbConfig DbConfig, maxAttempts int) (*gorm.DB, error) {
	attempt := 1
	for {
		db, success := connectToDb(dbConfig)
		if success {
			return db, nil
		}

		attempt++
		if attempt > maxAttempts {
			return nil, errors.New("max connection attempts reached")
		}

		fmt.Println("Retrying in 1 second...")
		time.Sleep(1 * time.Second)
	}
}

func connectToDb(dbConfig DbConfig) (*gorm.DB, bool) {
	var db *gorm.DB
	var err error
	args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DbName, dbConfig.Password)
	db, err = gorm.Open("postgres", args)
	if err == nil {
		return autoMigrateDb(db), true
	}
	return nil, false
}

func autoMigrateDb(db *gorm.DB) *gorm.DB {
	// AutoMigrate both Vocabulary and VocabularyCategory models
	db.AutoMigrate(&VocabularyGormRepository.Vocabulary{})
	db.AutoMigrate(&VocabularyGormRepository.Category{})

	// Create a migration to define the join table with foreign keys
	db.Table("vocabulary_categories").AddForeignKey("vocabulary_id", "vocabularies(id)", "CASCADE", "CASCADE")
	db.Table("vocabulary_categories").AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")

	return db
}
