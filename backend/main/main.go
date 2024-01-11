package main

import (
	"fmt"
	"os"
	"vocabulary/gormRepository"
	"vocabulary/main/util"
	"vocabulary/services"
)

func main() {
	config, err := util.LoadConfig("conf.json")
	if err != nil {
		fmt.Println("Config file error. Exiting.")
		os.Exit(1)
	}

	db, errDb := gormRepository.ConnectToDbWithMaxAttempts(config.DbConfig, 30)
	if errDb != nil {
		fmt.Println("DB: Max connection attempts reached. Exiting.")
		os.Exit(1)
	}

	gormTxRepositoryHandler := gormRepository.NewGormTxRepositoryHandler(db)

	endpoints := services.NewEndpoints(gormTxRepositoryHandler)
	endpoints.ListenAndServe(config.ApiPort)
}
