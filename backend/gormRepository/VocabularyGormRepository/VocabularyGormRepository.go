package VocabularyGormRepository

import (
	"github.com/jinzhu/gorm"
	"time"
	"vocabulary/entities/VocabularyEntity"
)

type Repository struct {
	tx *gorm.DB
}
type Vocabulary struct {
	Id           *uint `gorm:"primaryKey"`
	Words        *string
	Translation  *string
	UsedInPhrase *string
	Explanation  *string
	Categories   []*Category `gorm:"many2many:vocabulary_categories;"`
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
}

func (o Vocabulary) mapVocabularyToEntity() VocabularyEntity.Vocabulary {
	return VocabularyEntity.Vocabulary{
		Id:           o.Id,
		Words:        o.Words,
		Translation:  o.Translation,
		UsedInPhrase: o.UsedInPhrase,
		Explanation:  o.Explanation,
		Categories:   mapCategoriesToEntities(o.Categories),
		CreatedAt:    o.CreatedAt,
		UpdatedAt:    o.UpdatedAt,
	}
}

type Category struct {
	Id           *uint         `gorm:"primaryKey" json:"id"`
	Name         *string       `json:"name" gorm:"unique;not null"`
	Vocabularies []*Vocabulary `gorm:"many2many:vocabulary_categories;"`
	CreatedAt    *time.Time    `json:"createdAt"`
	UpdatedAt    *time.Time    `json:"updatedAt"`
}

func (o Repository) CreateVocabulary(vocabulary VocabularyEntity.Vocabulary) VocabularyEntity.Vocabulary {
	gormVocabulary := Vocabulary{
		Words:        vocabulary.Words,
		Translation:  vocabulary.Translation,
		UsedInPhrase: vocabulary.UsedInPhrase,
		Explanation:  vocabulary.Explanation,
	}
	o.tx.Create(&gormVocabulary)
	vocabulary.Id = gormVocabulary.Id
	vocabulary.UpdatedAt = gormVocabulary.UpdatedAt
	vocabulary.CreatedAt = gormVocabulary.CreatedAt
	return vocabulary
}

func (o Repository) UpdateVocabulary(vocabulary VocabularyEntity.Vocabulary) VocabularyEntity.Vocabulary {
	gormVocabulary := Vocabulary{
		Id:           vocabulary.Id,
		Words:        vocabulary.Words,
		Translation:  vocabulary.Translation,
		UsedInPhrase: vocabulary.UsedInPhrase,
		Explanation:  vocabulary.Explanation,
	}
	o.tx.Update(&gormVocabulary)
	vocabulary.UpdatedAt = gormVocabulary.UpdatedAt
	return vocabulary
}

func (o Repository) UpdateVocabularyWithCategories(vocabulary VocabularyEntity.Vocabulary, categories []string) {
	var dbVocabulary Vocabulary
	o.tx.Model(&Vocabulary{}).First(&dbVocabulary, vocabulary.Id)
	if dbVocabulary.Id != nil {
		gormVocabulary := Vocabulary{
			Id:           vocabulary.Id,
			Words:        vocabulary.Words,
			Translation:  vocabulary.Translation,
			UsedInPhrase: vocabulary.UsedInPhrase,
			Explanation:  vocabulary.Explanation,
		}
		o.tx.Update(&gormVocabulary)

		o.tx.Model(vocabulary).Association("Categories").Clear()

		for _, categoryName := range categories {
			dbCategory := Category{}
			if err := o.tx.Where("name = ?", categoryName).First(&dbCategory).Error; err != nil {
				dbCategory = Category{Name: &categoryName}
				o.tx.Create(&dbCategory)
			}
			o.tx.Model(vocabulary).Association("Categories").Append(dbCategory)
		}
	}
}

func (o Repository) GetAllVocabulariesWithCategories() []VocabularyEntity.Vocabulary {
	var dbVocabularies []Vocabulary
	err := o.tx.Model(&Vocabulary{}).Order("created_at desc").Preload("Categories").Find(&dbVocabularies).Error
	if err != nil {
		panic(err)
	}

	var vocabularies []VocabularyEntity.Vocabulary
	vocabularies = mapVocabulariesToEntities(dbVocabularies, vocabularies)

	return vocabularies
}

func (o Repository) FindVocabularyById(id uint) VocabularyEntity.Vocabulary {
	var dbVocabulary Vocabulary
	o.tx.Model(&Vocabulary{}).Preload("Categories").First(&dbVocabulary, id)
	return dbVocabulary.mapVocabularyToEntity()
}

func (o Repository) FindCategories() []VocabularyEntity.Category {
	var dbCategories []*Category
	o.tx.Order("created_at DESC").Find(&dbCategories)
	return mapCategoriesToEntities(dbCategories)
}

func mapVocabulariesToEntities(dbVocabularies []Vocabulary, vocabularies []VocabularyEntity.Vocabulary) []VocabularyEntity.Vocabulary {
	for _, dbVocabulary := range dbVocabularies {
		vocabularies = append(vocabularies, VocabularyEntity.Vocabulary{
			Id:           dbVocabulary.Id,
			Words:        dbVocabulary.Words,
			Translation:  dbVocabulary.Translation,
			UsedInPhrase: dbVocabulary.UsedInPhrase,
			Explanation:  dbVocabulary.Explanation,
			Categories:   mapCategoriesToEntities(dbVocabulary.Categories),
			CreatedAt:    dbVocabulary.CreatedAt,
			UpdatedAt:    dbVocabulary.UpdatedAt,
		})
	}
	return vocabularies
}

func mapCategoriesToEntities(dcCategories []*Category) []VocabularyEntity.Category {
	var categories []VocabularyEntity.Category
	for _, dbCategory := range dcCategories {
		categories = append(categories, VocabularyEntity.Category{
			Id:        dbCategory.Id,
			Name:      dbCategory.Name,
			CreatedAt: dbCategory.CreatedAt,
			UpdatedAt: dbCategory.UpdatedAt,
		})
	}

	return categories
}

func New(tx *gorm.DB) Repository {
	//tx.LogMode(true)
	return Repository{
		tx: tx,
	}
}
