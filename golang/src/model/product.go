package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string
	Price      float64
	Categories []Category `gorm:"many2many:product_categories;"`
}
