package main

import (
	"os"
	"vocabulary/gormRepository"
	"vocabulary/logger"
	"vocabulary/main/util"
	"vocabulary/services/vocabularyEndpoints"
)

func main() {
	logger.InitLogger()

	config, err := util.LoadConfig("conf.json")
	if err != nil {
		logger.LogInfo("Config file error. Exiting.")
		os.Exit(1)
	}

	db, err := gormRepository.ConnectToDbWithMaxAttempts(config.DbConfig, 30)
	if err != nil {
		logger.LogInfo("DB: Max connection attempts reached. Exiting.")
		os.Exit(1)
	}
	gormTxRepositoryHandler := gormRepository.NewGormTxRepositoryHandler(db)

	endpoints := vocabularyEndpoints.NewEndpoints(gormTxRepositoryHandler)
	endpoints.ListenAndServe(config.ApiPort)
}
