package service

import (
	"golang/src/model"
	"golang/src/repository"
)

type CategoryService interface {
	Create(*model.Category) (*model.Category, error)
	GetByID(id float64) (*model.Category, error)
	GetAll() ([]model.Category, error)
	Update(id float64, category *model.Category) (*model.Category, error)
	Delete(id float64) error
}

type categoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository: categoryRepository}
}

func (service categoryService) Create(category *model.Category) (*model.Category, error) {
	return service.categoryRepository.Create(category)
}

func (service categoryService) GetByID(id float64) (*model.Category, error) {
	return service.categoryRepository.GetByID(id)
}

func (service categoryService) GetAll() ([]model.Category, error) {
	return service.categoryRepository.GetAll()
}

func (service categoryService) Update(id float64, category *model.Category) (*model.Category, error) {
	return service.categoryRepository.Update(id, category)
}

func (service categoryService) Delete(id float64) error {
	return service.categoryRepository.Delete(id)
}
