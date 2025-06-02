package service

import (
	"golang/src/model"
	"golang/src/repository"
)

type ProductService interface {
	Create(*model.Product) (*model.Product, error)
	GetByID(id float64) (*model.Product, error)
	GetAll() ([]model.Product, error)
	Update(id float64, product *model.Product) (*model.Product, error)
	Delete(id float64) error
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	return &productService{productRepository: productRepository}
}

type productService struct {
	productRepository repository.ProductRepository
}

func (service productService) Create(product *model.Product) (*model.Product, error) {
	return service.productRepository.Create(product)
}

func (service productService) GetByID(id float64) (*model.Product, error) {
	return service.productRepository.GetByID(id)
}

func (service productService) GetAll() ([]model.Product, error) {
	return service.productRepository.GetAll()
}

func (service productService) Update(id float64, product *model.Product) (*model.Product, error) {
	return service.productRepository.Update(id, product)
}

func (service productService) Delete(id float64) error {
	return service.productRepository.Delete(id)
}
