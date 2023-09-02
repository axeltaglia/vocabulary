package services

import "vocabulary/modules"

type VocabularyWithCategories struct {
	Vocabulary modules.Vocabulary `json:"vocabulary"`
	Categories []string           `json:"categories"`
}
