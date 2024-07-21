package routes

import (
	"back-end/controllers"
	"back-end/middleware"

	"github.com/gin-gonic/gin"
)

type GoodsRouteController struct {
	userController controllers.GoodsController
}

func NewRouteGoodsController(goodsController controllers.GoodsController) GoodsRouteController {
	return GoodsRouteController{goodsController}
}

func (uc *GoodsRouteController) GoodsRoute(rg *gin.RouterGroup) {

	router := rg.Group("goods")
	router.GET("/:id", middleware.DeserializeUser(), uc.userController.GetById)
	router.PUT("/:id/img", middleware.DeserializeUser(), uc.userController.GetById)
	router.PUT("/:id/ref", middleware.DeserializeUser(), uc.userController.GetById)
	router.POST("/InsertGoods", middleware.DeserializeUser(), uc.userController.InsertGoods)
}
