package model

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	Amount float64
}
