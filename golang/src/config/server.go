package config

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang/src/controller"
	"golang/src/repository"
	"golang/src/service"
)

func Run() {
	ConnectDatabase()
	router := gin.New()
	router.Use(cors.New(GetConfig()))
	SetUpRouter(router)
	router.Run("localhost:8080")
}

func GetConfig() cors.Config {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	config.AllowCredentials = true
	return config
}

func SetUp() (controller.ProductController, controller.CategoryController, controller.BasketController, controller.PaymentController) {
	productRepo := repository.NewProductRepository(GormDB)
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	categoryRepo := repository.NewCategoryRepository(GormDB)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryController(categoryService)

	paymentRepository := repository.NewPaymentRepository(GormDB)
	paymentController := controller.NewPaymentController(paymentRepository)

	basketController := controller.NewBasketController(GormDB)

	return *productController, *categoryController, *basketController, *paymentController
}

func SetUpRouter(router *gin.Engine) *gin.Engine {

	productsController, categoryController, basketController, paymentController := SetUp()

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

	router.POST("/payments", paymentController.CreatePayment)

	return router
}
