package config

import (
	"github.com/gin-gonic/gin"
	"golang/src/controller"
	"golang/src/repository"
	"golang/src/service"
)

func Run() {
	ConnectDatabase()
	router := SetUpRouter()
	router.Run("localhost:8080")
}

func SetUp() (controller.ProductController, controller.CategoryController, controller.BasketController) {
	productRepo := repository.NewProductRepository(GormDB)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	categoryRepo := repository.NewCategoryRepository(GormDB)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryController(categoryService)

	basketController := controller.NewBasketController(GormDB)

	return *productController, *categoryController, *basketController
}

func SetUpRouter() *gin.Engine {
	router := gin.New()
	productsController, categoryController, basketController := SetUp()

	router.POST("/products", productsController.Create)
	router.GET("/products", productsController.GetAll)
	router.GET("/products/:id", productsController.GetByID)
	router.PUT("/products/:id", productsController.Update)
	router.DELETE("/products/:id", productsController.Delete)

	router.POST("/categories", categoryController.Create)
	router.GET("/categories", categoryController.GetAll)
	router.GET("/categories/:id", categoryController.GetByID)
	router.PUT("/categories/:id", categoryController.Update)
	router.DELETE("/categories/:id", categoryController.Delete)

	router.POST("/baskets", basketController.Create)
	router.GET("/baskets", basketController.GetAll)
	router.GET("/baskets/:id", basketController.GetByID)
	router.PUT("/baskets/:id", basketController.Update)
	router.DELETE("/baskets/:id", basketController.Delete)

	return router
}
