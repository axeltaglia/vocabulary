package vocabularyEndpoints

import "vocabulary/entities/VocabularyEntity"

type Vocabulary struct {
	Id           *uint      `json:"id"`
	Words        *string    `json:"words"`
	Translation  *string    `json:"translation"`
	UsedInPhrase *string    `json:"usedInPhrase"`
	Explanation  *string    `json:"explanation"`
	Categories   []Category `json:"categories"`
}

type Category struct {
	Id   *uint   `json:"id"`
	Name *string `json:"name"`
}

func (o *Vocabulary) MapFromEntity(vocabulary *VocabularyEntity.Vocabulary) {
	o.Id = vocabulary.Id
	o.Words = vocabulary.Words
	o.Translation = vocabulary.Translation
	o.UsedInPhrase = vocabulary.UsedInPhrase
	o.Explanation = vocabulary.Explanation
}
