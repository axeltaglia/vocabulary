package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"vocabulary/gormRepository"
	"vocabulary/main/util"
	"vocabulary/services"
)

// docker run --name pg_vocabulary_ctn -e POSTGRES_USER=vocabulary -e POSTGRES_PASSWORD=vocabulary -e POSTGRES_DB=vocabulary -e PGPORT=5435 -p 5435:5435 -v dbData:/var/lib/postgresql/data -d postgres
func main() {
	config, err := util.LoadConfig("conf.json")
	if err != nil {
		panic(err)
	}
	db := gormRepository.InitDb(config.DbConfig)

	gormTxRepositoryHandler := gormRepository.NewGormTxRepositoryHandler(db)

	endpoints := services.NewEndpoints(config.ApiPort, gormTxRepositoryHandler)
	endpoints.ListenAndServe()
}
