package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"version1/model"
	"version1/utils/errmsg"
)

// AddModel 添加已经训练好的模型
func AddModel(ctx *gin.Context) {

	// 1.接收参数
	var modelzoo model.ModelZoo
	_ = ctx.ShouldBind(&modelzoo)

	// 2.业务处理
	code := model.CheckModel(modelzoo.ModelName, modelzoo.DatasetName)
	if code == errmsg.ERROR_MODEL_EXISTED {
		ctx.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   modelzoo,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	} else {
		code = model.AddModel(&modelzoo)
	}
	if code == errmsg.ERROR {
		ctx.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   modelzoo,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	}
	// 3.返回响应
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   modelzoo,
		"msg":    errmsg.GetErrMsg(code),
	})
}

// ModelZooList 模型列表展示
func ModelZooList(ctx *gin.Context) {

	//业务处理
	data, code := model.ModelZooList()
	if code == errmsg.ERROR {
		ctx.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	}
	//返回响应
	ctx.HTML(http.StatusOK, "modelzoo.html", gin.H{
		"modelzoolist": data,
	})
}
