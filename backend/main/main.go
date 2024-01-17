package main

import (
	"fmt"
	"os"
	"vocabulary/gormRepository"
	"vocabulary/logger"
	"vocabulary/main/util"
	"vocabulary/services/vocabularyEndpoints"
)

func main() {
	config, err := util.LoadConfig("conf.json")
	if err != nil {
		fmt.Println("Config file error. Exiting.")
		os.Exit(1)
	}

	logger.InitLogger()
	logger.Log().Infof("holaaa")

	db, errDb := gormRepository.ConnectToDbWithMaxAttempts(config.DbConfig, 30)
	if errDb != nil {
		fmt.Println("DB: Max connection attempts reached. Exiting.")
		os.Exit(1)
	}

	gormTxRepositoryHandler := gormRepository.NewGormTxRepositoryHandler(db)

	endpoints := vocabularyEndpoints.NewEndpoints(gormTxRepositoryHandler)
	endpoints.ListenAndServe(config.ApiPort)
}
