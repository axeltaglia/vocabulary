package main

import (
	"os"

	"vocabulary/gormRepository"
	"vocabulary/logger"
	"vocabulary/logrusLogger"
	"vocabulary/main/util"
	"vocabulary/services/vocabularyEndpoints"
)

func main() {
	// Initialize logger
	logger.InitializeLogger(&logrusLogger.LogrusLogger{})

	// Load configuration
	config, err := util.LoadConfig("conf.json")
	if err != nil {
		logger.GetLogger().LogInfo("Config file error. Exiting.")
		os.Exit(1)
	}

	// Connect to the database
	db, err := gormRepository.ConnectToDbWithMaxAttempts(config.DbConfig, 5)
	if err != nil {
		logger.GetLogger().LogInfo("DB: Max connection attempts reached. Exiting.")
		os.Exit(1)
	}

	// Initialize GormTxRepositoryHandler
	txRepositoryHandler := gormRepository.NewGormTxRepositoryHandler(db)

	// Initialize endpoints
	endpoints := vocabularyEndpoints.NewEndpoints(txRepositoryHandler)

	// Start server
	endpoints.ListenAndServe(config.ApiPort)
}
