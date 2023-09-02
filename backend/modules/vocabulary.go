package modules

import (
	"time"
)

type Vocabulary struct {
	Id           uint       `gorm:"primaryKey" json:"id"`
	Words        string     `json:"words"`
	Translation  string     `json:"translation"`
	UsedInPhrase string     `json:"usedInPhrase"`
	Explanation  string     `json:"explanation"`
	Categories   []Category `json:"categories" gorm:"many2many:vocabulary_categories;"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    time.Time  `json:"updatedAt"`
}

type Category struct {
	Id           uint         `gorm:"primaryKey" json:"id"`
	Name         string       `json:"name" gorm:"unique;not null"`
	Vocabularies []Vocabulary `gorm:"many2many:vocabulary_categories;"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`
}
