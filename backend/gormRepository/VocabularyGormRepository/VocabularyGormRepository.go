package VocabularyGormRepository

import (
	"errors"
	"fmt"
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

func (o Repository) CreateVocabularyWithCategories(vocabulary *VocabularyEntity.Vocabulary, categories []string) (*VocabularyEntity.Vocabulary, error) {
	createVocabulary, _ := o.CreateVocabulary(vocabulary)
	gormVocabulary := mapVocabularyToDbVocabulary(createVocabulary)

	if err := o.tx.Model(gormVocabulary).Association("Categories").Error; err != nil {
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

	createdVocabulary := mapDbVocabularyToVocabulary(*gormVocabulary)
	return &createdVocabulary, nil
}

func (o Repository) UpdateVocabulary(vocabulary *VocabularyEntity.Vocabulary) (*VocabularyEntity.Vocabulary, error) {
	gormVocabulary := mapVocabularyToDbVocabulary(vocabulary)
	if err := o.tx.Save(gormVocabulary).Error; err != nil {
		logger.LogError("[Gorm Repository] Vocabulary couldn't be updated", err)
		return nil, err
	}

	vocabulary.UpdatedAt = gormVocabulary.UpdatedAt
	return vocabulary, nil
}

func (o Repository) DisassociateCategoriesFromVocabulary(vocabulary *VocabularyEntity.Vocabulary) error {
	gormVocabulary := mapVocabularyToDbVocabulary(vocabulary)
	if err := o.tx.Model(gormVocabulary).Association("Categories").Clear().Error; err != nil {
		logger.LogError("[Gorm Repository] Categories couldn't be disassociated to Vocabulary", err)
		return err
	}
	return nil
}

func (o Repository) AssociateCategoryToVocabulary(vocabulary *VocabularyEntity.Vocabulary, category *VocabularyEntity.Category) (*VocabularyEntity.Vocabulary, error) {
	gormVocabulary := mapVocabularyToDbVocabulary(vocabulary)
	gormCategory := mapCategoryToDbCategory(category)
	if err := o.tx.Model(gormVocabulary).Association("Categories").Append(gormCategory).Error; err != nil {
		logger.LogError("[Gorm Repository] Category couldn't be associated to Vocabulary", err)
		return nil, err
	}

	updatedVocabulary := mapDbVocabularyToVocabulary(*gormVocabulary)
	return &updatedVocabulary, nil
}

func (o Repository) CreateCategoryIfNotExist(categoryName string) (*VocabularyEntity.Category, error) {
	dbCategory := Category{}
	if err := o.tx.Where("name = ?", categoryName).First(&dbCategory).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		dbCategory = Category{Name: &categoryName}
		if err := o.tx.Create(&dbCategory).Error; err != nil {
			logger.LogError("[Gorm Repository] Category couldn't be created", err)
			return nil, err
		}
	}
	return mapDbCategoryToCategory(dbCategory), nil
}

func (o Repository) DeleteVocabularyById(id uint) error {
	if err := o.tx.Delete(&Vocabulary{}, id).Error; err != nil {
		logger.LogError(fmt.Sprintf("[Gorm Repository] Couldn't delete Vocabulary from id %d", id), err)
		return err
	}
	return nil
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
	var categories []*Category
	for _, categoryEntity := range vocabulary.Categories {
		categories = append(categories, &Category{
			Id:        categoryEntity.Id,
			Name:      categoryEntity.Name,
			CreatedAt: categoryEntity.CreatedAt,
			UpdatedAt: categoryEntity.UpdatedAt,
		})
	}
	return &Vocabulary{
		Id:           vocabulary.Id,
		Words:        vocabulary.Words,
		Translation:  vocabulary.Translation,
		UsedInPhrase: vocabulary.UsedInPhrase,
		Explanation:  vocabulary.Explanation,
		Categories:   categories,
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

func mapCategoryToDbCategory(category *VocabularyEntity.Category) *Category {
	return &Category{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func mapDbCategoryToCategory(category Category) *VocabularyEntity.Category {
	return &VocabularyEntity.Category{
		Id:        category.Id,
		Name:      category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}

func New(tx *gorm.DB) Repository {
	//tx.LogMode(true)
	return Repository{
		tx: tx,
	}
}
