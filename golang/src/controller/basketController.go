package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"golang/src/model"
	"gorm.io/gorm"
)

type BasketController struct {
	db *gorm.DB
}

func NewBasketController(db *gorm.DB) *BasketController {
	return &BasketController{db: db}
}

func (controller *BasketController) Create(c *gin.Context) {
	var basket model.Basket
	if err := c.ShouldBindJSON(&basket); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := controller.db.Create(&basket).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, basket)
}

func (controller *BasketController) GetByID(c *gin.Context) {
	id := c.Param("id")
	var basket model.Basket
	if err := controller.db.Preload("Products").First(&basket, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{"error": "Basket not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(200, basket)
}

func (controller *BasketController) GetAll(c *gin.Context) {
	var baskets []model.Basket
	if err := controller.db.Preload("Products").Find(&baskets).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, baskets)
}

func (controller *BasketController) Update(c *gin.Context) {
	id := c.Param("id")
	var basket model.Basket
	if err := c.ShouldBindJSON(&basket); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := controller.db.Model(&model.Basket{}).Where("id = ?", id).Updates(basket).Error; err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		return
	}
	var updatedBasket model.Basket
	if err := controller.db.First(&updatedBasket, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{"error": "Basket not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(200, updatedBasket)
}

func (controller *BasketController) Delete(c *gin.Context) {
	id := c.Param("id")
	if err := controller.db.Delete(&model.Basket{}, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(404, gin.H{"error": "Basket not found"})
		} else {
			c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
		}
		return
	}
}
