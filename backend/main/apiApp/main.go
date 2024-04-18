package main

import (
	"os"
	"vocabulary/gormRepository"
	"vocabulary/logger"
	"vocabulary/main/util"
	"vocabulary/services/vocabularyEndpoints"
	"vocabulary/slogJsonLogger"
)

func main() {
	// Initialize logger
	logger.InitializeLogger(&slogJsonLogger.SlogJsonLogger{})

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
	if err := endpoints.ListenAndServe(config.ApiPort); err != nil {
		logger.GetLogger().LogError("server coundn't start", err)
		os.Exit(1)
	}
}
