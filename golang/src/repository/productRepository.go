package repository

import (
	"golang/src/model"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *model.Product) (*model.Product, error)
	GetByID(id float64) (*model.Product, error)
	GetAll() ([]model.Product, error)
	Update(id float64, product *model.Product) (*model.Product, error)
	Delete(id float64) error
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (repository productRepository) Create(product *model.Product) (*model.Product, error) {
	if err := repository.db.Create(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repository productRepository) GetByID(id float64) (*model.Product, error) {
	var product model.Product
	if err := repository.db.Preload("Categories").First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (repository productRepository) GetAll() ([]model.Product, error) {
	var products []model.Product
	if err := repository.db.Preload("Categories").Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (repository productRepository) Update(id float64, product *model.Product) (*model.Product, error) {
	if err := repository.db.Model(&model.Product{}).Where("id = ?", id).Updates(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (repository productRepository) Delete(id float64) error {
	if err := repository.db.Delete(&model.Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
