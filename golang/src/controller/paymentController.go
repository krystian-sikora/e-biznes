package controller

import (
	"github.com/gin-gonic/gin"
	"golang/src/model"
	"golang/src/repository"
)

type PaymentController struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentController(paymentRepository repository.PaymentRepository) *PaymentController {
	return &PaymentController{paymentRepository: paymentRepository}
}

func (pc *PaymentController) CreatePayment(c *gin.Context) {
	var payment model.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid payment data"})
		return
	}

	createdPayment, err := pc.paymentRepository.Create(&payment)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Failed to create payment"})
		return
	}
	c.JSON(201, createdPayment)
}
