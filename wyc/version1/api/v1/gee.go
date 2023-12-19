package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"version1_copy/model"
	"version1_copy/utils/errmsg"
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
	code, imgurl := model.GetImage(&geeparam)
	ctx.JSON(http.StatusOK, gin.H{
		"status":   code,
		"msg":      errmsg.GetErrMsg(code),
		"imageurl": imgurl,
	})
}

func AddGee(ctx *gin.Context) {
	// 1.接收参数
	var gee model.Gee
	isoverwrite := ctx.PostForm("isoverwrite")
	_ = ctx.ShouldBind(&gee)
	gee.Uid = 1

	//2. 写入数据库
	if isoverwrite == "0" {
		code := model.CheckImageUserdb(gee.Image_Name, gee.Uid)
		// 名称重复，返回选择是否覆盖
		if code == errmsg.ERROR_IMAGE_EXIT {
			ctx.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errmsg.GetErrMsg(code),
			})
			return
		} else {
			// 该名称的图像不存在，直接写入数据库
			gee.Save_Path = strings.Split(gee.Save_Path, "static/")[0]
			code = model.AddImage(gee.Image_Name, gee.Save_Path+image_save_path+gee.Image_Name)
			if code != errmsg.SUCCESS {
				ctx.JSON(http.StatusOK, gin.H{
					"status": "保存失败",
					"msg":    errmsg.GetErrMsg(code),
				})
				return
			}
			// 3. 返回响应
			ctx.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errmsg.GetErrMsg(code),
			})
			return
		}
	}

	// 该名称的图像存在，且用户选择覆盖，更新数据库
	gee.Save_Path = strings.Split(gee.Save_Path, gee.Image_Name)[0]
	code := model.UpdateImageUserdb(gee.Image_Name, gee.Save_Path+gee.Image_Name)
	// 3. 返回响应
	if code != errmsg.SUCCESS {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "保存失败",
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
