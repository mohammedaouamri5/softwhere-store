package main

import (
	"log"
	"net/http"

	"back-end/controllers"
	"back-end/initializers"
	"back-end/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	server              *gin.Engine
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController

	GoodsController      controllers.GoodsController
	GoodsRouteController routes.GoodsRouteController

	TransactionController      controllers.TransactionController
	TransactionRouteController routes.TransactionRouteController


)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	GoodsController = controllers.NewGoodsController(initializers.DB)
	GoodsRouteController = routes.NewRouteGoodsController(GoodsController)

	TransactionController = controllers.NewTransactionController(initializers.DB)
	TransactionRouteController = routes.NewRouteTransactionController(TransactionController)

	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin, "*"}
	corsConfig.AllowCredentials = true

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	GoodsRouteController.GoodsRoute(router)
	TransactionRouteController.TransactionRoute(router)
	// PostRouteController.PostRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))
}
