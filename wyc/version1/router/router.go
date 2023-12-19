package router

import (
	"github.com/gin-gonic/gin"
	"version1_copy/api/v1"
	"version1_copy/utils"
)

func SetRouter() *gin.Engine {

	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")
	router := r.Group("/api/v1")
	{
		// 用户模块路由接口
		router.GET("login", v1.Login)
		//router.POST("login", v1.LoginCheck)
		//router.GET("register", v1.Register)
		router.POST("register", v1.RegisterCheck)
		router.POST("home", v1.Home)
		router.GET("index", v1.Index)

		router.POST("modelzoo/add", v1.AddModel)
		router.GET("modelzoo", v1.ModelZooList)
		router.GET("train", v1.Train)
		router.GET("inference", v1.InferencePage)
		router.POST("inference", v1.Inference)

		router.GET("gee", v1.GeePage)
		router.POST("gee", v1.Gee)
		router.POST("addgee", v1.AddGee)

		//用户数据库
		router.GET("userdb", v1.UserdbPage)
		router.GET("getuserdbdata", v1.GetUserDbData)

		router.POST("trans", v1.Transpose)

		router.GET("test", v1.TestPage)
		router.POST("test", v1.Test)
	}
	return r
}
