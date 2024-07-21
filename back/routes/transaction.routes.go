package routes

import (
	"net/http"

	"back-end/controllers"
	// "back-end/middleware"

	"github.com/gin-gonic/gin"
)

type TransactionRouteController struct {
	postController controllers.TransactionController
}

func NewRouteTransactionController(postController controllers.TransactionController) TransactionRouteController {
	return TransactionRouteController{postController}
}

func (uc *TransactionRouteController) TransactionRoute(rg *gin.RouterGroup) {

	router := rg.Group("transaction")
	router.GET("/me",
		
		func(ctx *gin.Context) {
			message := "Welcome to Golang with Gorm and Postgres"
			ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
		},
	)
}

// func (pc *PostRouteController) PostRoute(rg *gin.RouterGroup) {
// 	router := rg.Group("posts")
// 	router.Use(middleware.DeserializeUser())
// 	router.POST("/", pc.postController.CreatePost)
// 	router.GET("/", pc.postController.FindPosts)
// 	router.PUT("/:postId", pc.postController.UpdatePost)
// 	router.GET("/:postId", pc.postController.FindPostById)
// 	router.DELETE("/:postId", pc.postController.DeletePost)
// }
