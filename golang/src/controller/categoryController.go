package controller

import (
	"github.com/gin-gonic/gin"
	"golang/src/model"
	"golang/src/service"
	"net/http"
	"strconv"
)

type CategoryController struct {
	categoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{categoryService: categoryService}
}

const InvalidIdFormat = "Invalid ID format"

func (controller *CategoryController) Create(c *gin.Context) {
	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := controller.categoryService.Create(&category)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdCategory)
}

func (controller *CategoryController) GetByID(c *gin.Context) {
	id := c.Param("id")
	floatId, err := strconv.ParseFloat(id, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": InvalidIdFormat})
		return
	}

	category, err := controller.categoryService.GetByID(floatId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if category == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (controller *CategoryController) GetAll(c *gin.Context) {
	categories, err := controller.categoryService.GetAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (controller *CategoryController) Update(c *gin.Context) {
	id := c.Param("id")
	floatId, err := strconv.ParseFloat(id, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": InvalidIdFormat})
		return
	}

	var category model.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory, err := controller.categoryService.Update(floatId, &category)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedCategory)
}

func (controller *CategoryController) Delete(c *gin.Context) {
	id := c.Param("id")
	floatId, err := strconv.ParseFloat(id, 64)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": InvalidIdFormat})
		return
	}

	if err := controller.categoryService.Delete(floatId); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
