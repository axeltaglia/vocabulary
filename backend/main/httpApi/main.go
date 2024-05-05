package main

import (
	"fmt"
	"os"
	"vocabulary/logger"
	"vocabulary/main/util"
	"vocabulary/slogJsonLogger"
	"vocabulary/sqlRepository"
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
	db, err := sqlRepository.ConnectToDbWithMaxAttempts(sqlRepository.DbConfig{
		Host:     config.DbConfig.Host,
		Port:     config.DbConfig.Port,
		User:     config.DbConfig.User,
		DbName:   config.DbConfig.DbName,
		Password: config.DbConfig.Password,
	}, 5)
	if err != nil {
		logger.GetLogger().LogInfo("DB: Max connection attempts reached. Exiting.")
		os.Exit(1)
	}

	// Initialize GormTxRepositoryHandler
	txRepositoryHandler := sqlRepository.NewSqlTxRepositoryHandler(db)

	apiServer := NewApiServer(":8080", &txRepositoryHandler)

	apiServer.HandleEndpoints()

	fmt.Printf("Server started. Listening in port %s\n", config.DbConfig.Port)

	if err := apiServer.ListenAndServe(); err != nil {
		logger.GetLogger().LogError("server coundn't start", err)
		os.Exit(1)
	}
}
