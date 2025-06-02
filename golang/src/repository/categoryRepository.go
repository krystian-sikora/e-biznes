package repository

import (
	"golang/src/model"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *model.Category) (*model.Category, error)
	GetByID(id float64) (*model.Category, error)
	GetAll() ([]model.Category, error)
	Update(id float64, category *model.Category) (*model.Category, error)
	Delete(id float64) error
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (repository categoryRepository) Create(category *model.Category) (*model.Category, error) {
	if err := repository.db.Create(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (repository categoryRepository) GetByID(id float64) (*model.Category, error) {
	var category model.Category
	if err := repository.db.First(&category, id).Error; err != nil {
		return nil, err
	}
	return &category, nil
}

func (repository categoryRepository) GetAll() ([]model.Category, error) {
	var categories []model.Category
	if err := repository.db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func (repository categoryRepository) Update(id float64, category *model.Category) (*model.Category, error) {
	if err := repository.db.Model(&model.Category{}).Where("id = ?", id).Updates(category).Error; err != nil {
		return nil, err
	}
	return category, nil
}

func (repository categoryRepository) Delete(id float64) error {
	if err := repository.db.Delete(&model.Category{}, id).Error; err != nil {
		return err
	}
	return nil
}
