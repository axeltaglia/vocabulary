package VocabularyGormRepository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
	"vocabulary/entities/VocabularyEntity"
	"vocabulary/logger"
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

type Category struct {
	Id           *uint         `gorm:"primaryKey" json:"id"`
	Name         *string       `json:"name" gorm:"unique;not null"`
	Vocabularies []*Vocabulary `gorm:"many2many:vocabulary_categories;"`
	CreatedAt    *time.Time    `json:"createdAt"`
	UpdatedAt    *time.Time    `json:"updatedAt"`
}

func (o Repository) CreateVocabulary(vocabulary *VocabularyEntity.Vocabulary) (*VocabularyEntity.Vocabulary, error) {
	gormVocabulary := mapVocabularyToDbVocabulary(vocabulary)

	if err := o.tx.Create(&gormVocabulary).Error; err != nil {
		logger.LogError("DB: Vocabulary couldn't be created", err)
		return nil, err
	}
	vocabulary.Id = gormVocabulary.Id
	vocabulary.UpdatedAt = gormVocabulary.UpdatedAt
	vocabulary.CreatedAt = gormVocabulary.CreatedAt
	return vocabulary, nil
}

func (o Repository) UpdateVocabulary(vocabulary *VocabularyEntity.Vocabulary) (*VocabularyEntity.Vocabulary, error) {
	gormVocabulary := mapVocabularyToDbVocabulary(vocabulary)

	if err := o.tx.Update(&gormVocabulary).Error; err != nil {
		return nil, err
	}
	vocabulary.UpdatedAt = gormVocabulary.UpdatedAt
	return vocabulary, nil
}

func (o Repository) UpdateVocabularyWithCategories(vocabulary *VocabularyEntity.Vocabulary, categories []string) (*VocabularyEntity.Vocabulary, error) {
	gormVocabulary := mapVocabularyToDbVocabulary(vocabulary)
	if err := o.tx.Save(gormVocabulary).Error; err != nil {
		return nil, err
	}

	if err := o.tx.Model(gormVocabulary).Association("Categories").Clear().Error; err != nil {
		return nil, err
	}

	for _, categoryName := range categories {
		dbCategory := Category{}
		if err := o.tx.Where("name = ?", categoryName).First(&dbCategory).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			dbCategory = Category{Name: &categoryName}
			if err := o.tx.Create(&dbCategory).Error; err != nil {
				return nil, err
			}
		}
		if err := o.tx.Model(gormVocabulary).Association("Categories").Append(dbCategory).Error; err != nil {
			return nil, err
		}
	}

	updatedVocabulary := mapDbVocabularyToVocabulary(*gormVocabulary)
	return &updatedVocabulary, nil
}

func (o Repository) GetAllVocabulariesWithCategories() ([]VocabularyEntity.Vocabulary, error) {
	var dbVocabularies []Vocabulary
	if err := o.tx.Model(&Vocabulary{}).Order("created_at desc").Preload("Categories").Find(&dbVocabularies).Error; err != nil {
		logger.LogError(err.Error(), err)
		return nil, err
	}

	var vocabularies []VocabularyEntity.Vocabulary
	vocabularies = mapDbVocabulariesToVocabularies(dbVocabularies, vocabularies)

	return vocabularies, nil
}

func (o Repository) FindVocabularyById(id uint) (*VocabularyEntity.Vocabulary, error) {
	var dbVocabulary Vocabulary
	err := o.tx.Model(&Vocabulary{}).Preload("Categories").First(&dbVocabulary, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// Handle record not found error...
	}
	vocabulary := mapDbVocabularyToVocabulary(dbVocabulary)
	return &vocabulary, nil
}

func (o Repository) FindCategories() []VocabularyEntity.Category {
	var dbCategories []*Category
	o.tx.Order("created_at DESC").Find(&dbCategories)
	return mapDbCategoriesToCategories(dbCategories)
}

func mapDbVocabularyToVocabulary(dbVocabulary Vocabulary) VocabularyEntity.Vocabulary {
	return VocabularyEntity.Vocabulary{
		Id:           dbVocabulary.Id,
		Words:        dbVocabulary.Words,
		Translation:  dbVocabulary.Translation,
		UsedInPhrase: dbVocabulary.UsedInPhrase,
		Explanation:  dbVocabulary.Explanation,
		Categories:   mapDbCategoriesToCategories(dbVocabulary.Categories),
		CreatedAt:    dbVocabulary.CreatedAt,
		UpdatedAt:    dbVocabulary.UpdatedAt,
	}
}

func mapVocabularyToDbVocabulary(vocabulary *VocabularyEntity.Vocabulary) *Vocabulary {
	return &Vocabulary{
		Id:           vocabulary.Id,
		Words:        vocabulary.Words,
		Translation:  vocabulary.Translation,
		UsedInPhrase: vocabulary.UsedInPhrase,
		Explanation:  vocabulary.Explanation,
	}
}

func mapDbVocabulariesToVocabularies(dbVocabularies []Vocabulary, vocabularies []VocabularyEntity.Vocabulary) []VocabularyEntity.Vocabulary {
	for _, dbVocabulary := range dbVocabularies {
		vocabularies = append(vocabularies, VocabularyEntity.Vocabulary{
			Id:           dbVocabulary.Id,
			Words:        dbVocabulary.Words,
			Translation:  dbVocabulary.Translation,
			UsedInPhrase: dbVocabulary.UsedInPhrase,
			Explanation:  dbVocabulary.Explanation,
			Categories:   mapDbCategoriesToCategories(dbVocabulary.Categories),
			CreatedAt:    dbVocabulary.CreatedAt,
			UpdatedAt:    dbVocabulary.UpdatedAt,
		})
	}
	return vocabularies
}

func mapDbCategoriesToCategories(dcCategories []*Category) []VocabularyEntity.Category {
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
