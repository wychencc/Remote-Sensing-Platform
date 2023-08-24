package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"version1/model"
	"version1/utils/errmsg"
)

var type_name string

const (
	image_save_path  = "inference/images/"
	result_save_path = "inference/results/"
)

// Inference 渲染页面
func Inference(ctx *gin.Context) {

	model_name := ctx.PostForm("model_name")
	dataset_name := ctx.PostForm("dataset_name")

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": errmsg.ERROR,
			"data":   err,
			"msg":    errmsg.GetErrMsg(errmsg.ERROR),
		})
		return
	}

	image_name := file.Filename
	image_path := image_save_path + image_name
	_ = ctx.SaveUploadedFile(file, image_path)

	// 拿到模型存储路径
	model_path, code := model.GetModelPath(model_name, dataset_name)
	if code == errmsg.ERROR {
		ctx.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   model_path,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	}

	//根据原图保存路径和模型保存路径做推理
	result_path, result_name := model.GetResult(model_path, image_name, image_save_path, result_save_path)

	//此次推理记录到数据库
	code = model.AddInference(image_name, result_name, image_path, result_path)
	if code == errmsg.ERROR {
		ctx.JSON(http.StatusOK, gin.H{
			"status": code,
			"data":   model_path,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"image_name":  image_name,
		"image_path":  image_path,
		"result_name": result_name,
		"result_path": result_path,
	})
}

// InferencePage 执行推理任务
func InferencePage(ctx *gin.Context) {

	//获取选项数据的路由处理函数
	level := ctx.Query("level")
	value := ctx.Query("value")
	if level == "" {
		ctx.HTML(http.StatusOK, "inference.html", nil)
		return
	}
	if level == "second" {
		type_name = value
		data, _ := model.GetModelList(value)
		ctx.JSON(http.StatusOK, data)
		return
	}
	if level == "third" {
		data, _ := model.GetDataSetList(type_name, value)
		ctx.JSON(http.StatusOK, data)
		return
	}
	//data, code := model.GetModelList()
	//ctx.JSON(http.StatusOK, gin.H{
	//	"status": code,
	//	"data":   data,
	//	"msg":    errmsg.GetErrMsg(code),
	//})
}
