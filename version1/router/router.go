package router

import (
	"github.com/gin-gonic/gin"
	v1 "version1/api/v1"
	"version1/utils"
)

func SetRouter() {

	gin.SetMode(utils.AppMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	router := r.Group("/api/v1")
	{
		// 用户模块路由接口
		router.GET("login", v1.Login)
		router.POST("login", v1.LoginCheck)
		router.GET("register", v1.Register)
		router.POST("register", v1.RegisterCheck)
		router.GET("index", v1.Index)

		router.POST("modelzoo/add", v1.AddModel)
		router.GET("modelzoo", v1.ModelZooList)

		router.GET("/inference", v1.InferencePage)
		router.POST("/inference", v1.Inference)

		router.GET("/gee", v1.GeePage)
		router.POST("/gee", v1.Gee)

	}
	r.Run(utils.HttpPort)
}
