package repository

import (
	"golang/src/model"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *model.Payment) (*model.Payment, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db: db}
}

func (repository paymentRepository) Create(payment *model.Payment) (*model.Payment, error) {
	if err := repository.db.Create(payment).Error; err != nil {
		return nil, err
	}
	return payment, nil
}
