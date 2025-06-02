package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang/src/model"
	"golang/src/service"
	"net/http"
	"strconv"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService service.ProductService) *ProductController {
	return &ProductController{productService: productService}
}

func (pc *ProductController) Create(c *gin.Context) {
	var product model.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdProduct, err := pc.productService.Create(&product)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdProduct)
}

func (pc *ProductController) GetByID(c *gin.Context) {
	id := c.Param("id")
	floatId, err := strconv.ParseFloat(id, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	product, err := pc.productService.GetByID(floatId)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if product == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) GetAll(c *gin.Context) {
	products, err := pc.productService.GetAll()

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) Update(c *gin.Context) {
	id := c.Param("id")
	floatId, err := strconv.ParseFloat(id, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var product model.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProduct, err := pc.productService.Update(floatId, &product)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedProduct)
}

func (pc *ProductController) Delete(c *gin.Context) {
	id := c.Param("id")
	floatId, err := strconv.ParseFloat(id, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = pc.productService.Delete(floatId)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
