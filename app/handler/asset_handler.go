package handler

import (
	"assets-management/app/model"
	"assets-management/app/service"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(engine *gin.Engine) {
	engine.POST("/assets", calcAssets)
}

func calcAssets(ctx *gin.Context) {
	var req []model.AssetRequest
	if err := ctx.ShouldBind(&req); err != nil {
		fmt.Println(err.Error())
		return
	}
	var res = service.GetAssetsTotal(req)
	ctx.JSON(200, res)
}
