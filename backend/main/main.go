package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"os"
	"strings"
	"time"
	"vocabulary/modules"
	"vocabulary/services"
	"vocabulary/util"
)

// docker run --name pg_vocabulary_ctn -e POSTGRES_USER=vocabulary -e POSTGRES_PASSWORD=vocabulary -e POSTGRES_DB=vocabulary -e PGPORT=5435 -p 5435:5435 -v dbData:/var/lib/postgresql/data -d postgres
func main() {
	config, errConf := loadConfig("conf.json")
	util.CheckErr(errConf)
	db := initDb(config.DbConfig)
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	router.Use(cors.New(corsConfig))

	endpoints := services.NewEndpoints(router, db)
	endpoints.Handle()

	err := http.ListenAndServe(":"+config.ApiPort, router)
	util.CheckErr(err)
}

// Los valores de estas variables las tenes que leer del environment del sistema operativo.
type DbConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Config struct {
	ApiPort  string   `json:"apiPort"`
	DbConfig DbConfig `json:"dbConfig"`
}

func initDb(dbConfig DbConfig) *gorm.DB {
	maxAttempts := 30
	attempt := 1
	for {
		var db *gorm.DB
		var err error
		args := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Database, dbConfig.Password)
		db, err = gorm.Open("postgres", args)
		if err == nil {
			// AutoMigrate both Vocabulary and VocabularyCategory models
			db.AutoMigrate(&modules.Vocabulary{})
			db.AutoMigrate(&modules.Category{})

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

func loadConfig(configFileName string) (*Config, error) {
	// Read the env.json file
	file, err := os.ReadFile(configFileName)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into a map[string]interface{}
	var placeholders map[string]interface{}
	err = json.Unmarshal(file, &placeholders)
	if err != nil {
		return nil, err
	}

	// Replace placeholders with environment variables
	replacedPlaceholders := replacePlaceholders(placeholders)

	// Marshal the updated placeholders back to JSON
	updatedJSON, err := json.Marshal(replacedPlaceholders)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into the Config struct
	var config Config
	err = json.Unmarshal(updatedJSON, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func replacePlaceholders(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		for key, val := range v {
			v[key] = replacePlaceholders(val)
		}
	case string:
		if strings.HasPrefix(v, "{{.") && strings.HasSuffix(v, "}}") {
			envVar := strings.TrimPrefix(strings.TrimSuffix(v, "}}"), "{{.")
			replacedValue := os.Getenv(envVar)
			if replacedValue == "" {
				// If the environment variable is not set, keep the placeholder
				return v
			}
			return replacedValue
		}
	}
	return data
}
