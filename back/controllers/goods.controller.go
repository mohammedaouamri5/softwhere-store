package controllers

import ( 
	"net/http"

	"back-end/initializers"
	"back-end/models"
	"back-end/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type GoodsController struct {
	DB *gorm.DB
}

func NewGoodsController(DB *gorm.DB) GoodsController {
	return GoodsController{DB}   
}

func (uc *GoodsController) GetById(ctx *gin.Context) {
	id := ctx.Param("id") 
	var goods models.Goods; 
	
    query := "SELECT * FROM public.goods G WH ERE G.GOODS_ID=$1";

	if result  := initializers.DB.Raw(query, id).Scan(&goods); result.Error != nil {
 
		ctx.JSON(http.StatusInternalServerError ,  utils.ServerErrorResponse( result.Error))
		return 
	}
	
	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": goods })
}

func (uc *GoodsController) InsertGoods(ctx *gin.Context) {
	
	var goods models.Goods; 
	if err := ctx.ShouldBindJSON(&goods); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if goods.Photo == "" {
        goods.Photo = "/img/default/goods" // or any default value you prefer
    }


	var insertedGoods models.Goods
    query := "INSERT INTO public.goods (REF, Photo) VALUES ($1, $2) RETURNING GOODS_ID, REF, Photo"
    if err := initializers.DB.Raw(query, goods.REF, goods.Photo).Scan(&insertedGoods).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": insertedGoods})

}


