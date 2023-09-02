package modules

import "github.com/jinzhu/gorm"

type Vocabulary struct {
	gorm.Model
	Words        string `json:"words"`
	Translation  string `json:"translation"`
	UsedInPhrase string `json:"usedInPhrase"`
	Explanation  string `json:"explanation"`
}
