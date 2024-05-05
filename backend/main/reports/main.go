package main

import (
	"fmt"
	"os"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/gormRepository"
	"vocabulary/logger"
	"vocabulary/main/util"
)

func main() {
	// Initialize logger
	logger.InitializeLogger(&logger.LogrusLogger{})

	// Load configuration
	config, err := util.LoadConfig("conf.json")
	if err != nil {
		logger.GetLogger().LogInfo("Config file error. Exiting.")
		os.Exit(1)
	}

	// Connect to the database with maximum connection attempts
	db, err := gormRepository.ConnectToDbWithMaxAttempts(gormRepository.DbConfig{
		Host:     config.DbConfig.Host,
		Port:     config.DbConfig.Port,
		DbName:   config.DbConfig.DbName,
		Password: config.DbConfig.Password,
	}, 5)
	if err != nil {
		logger.GetLogger().LogInfo("DB: Max connection attempts reached. Exiting.")
		os.Exit(1)
	}

	// Initialize GormTxRepositoryHandler
	gormTxRepositoryHandler := gormRepository.NewGormTxRepositoryHandler(db)

	// Initialize VocabularyEntity
	vocabularyRepository := gormTxRepositoryHandler.GetTxRepositoryFactory().GetVocabularyRepository()
	vocabularyEntity := VocabularyEntity.New(vocabularyRepository)

	// Print report
	err = printAmountCategoriesFrequencyReport(vocabularyEntity)
	if err != nil {
		logger.GetLogger().LogError("Report failed.", err)
		os.Exit(1)
	}
}

func printAmountCategoriesFrequencyReport(vocabularyEntity VocabularyEntity.Entity) error {
	// Get all vocabularies with categories
	vocabularies, err := vocabularyEntity.GetAllVocabulariesWithCategories()
	if err != nil {
		logger.GetLogger().LogError("Error", err)
		return err
	}

	// Amount of categories frequency
	categoriesMap := make(map[int]int)
	for _, vocabulary := range vocabularies {
		categoriesMap[len(vocabulary.Categories)]++
	}

	// Print outout
	for i := 0; i < 5; i++ {
		fmt.Printf("There are %d words that have %d categories.\n", categoriesMap[i], i)
	}

	return nil
}
