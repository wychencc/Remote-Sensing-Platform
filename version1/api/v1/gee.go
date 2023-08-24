package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"version1/model"
)

// GeePage 渲染页面
func GeePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "gee.html", nil)
}

// Gee 执行Gee业务
func Gee(ctx *gin.Context) {

	var geeparam model.GeeParam
	//1. 接收参数
	_ = ctx.ShouldBind(&geeparam)
	ctx.JSON(http.StatusOK, gin.H{
		"data":       geeparam,
		"start_date": geeparam.Start_date,
		"end_date":   geeparam.End_date,
	})

	model.GetImage(&geeparam)
}
