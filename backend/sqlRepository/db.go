package sqlRepository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"time"
	"vocabulary/logger"
)

type DbConfig struct {
	Host     string
	Port     string
	User     string
	DbName   string
	Password string
}

func ConnectToDbWithMaxAttempts(dbConfig DbConfig, maxAttempts int) (*sql.DB, error) {
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

func connectToDb(dbConfig DbConfig) (*sql.DB, bool) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.DbName, dbConfig.Password))
	if err != nil {
		logger.GetLogger().LogInfo("couldn't connect to the db")
		return nil, false
	}
	return db, true
}
