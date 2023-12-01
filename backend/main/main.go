package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"vocabulary/gormRepository"
	"vocabulary/main/util"
	"vocabulary/services"
)

func main() {
	config, err := util.LoadConfig("conf.json")
	if err != nil {
		panic(err)
	}

	db := gormRepository.InitDb(config.DbConfig)
	gormTxRepositoryHandler := gormRepository.NewGormTxRepositoryHandler(db)

	endpoints := services.NewEndpoints(gormTxRepositoryHandler)
	endpoints.ListenAndServe(config.ApiPort)
}
