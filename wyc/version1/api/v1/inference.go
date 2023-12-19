package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"version1_copy/model"
	"version1_copy/utils/errmsg"
)

//var type_name string

const (
	base_path        = "http://localhost:6060/"
	image_save_path  = "static/inference/"
	image_tans_path  = "static/trans/"
	result_save_path = "static/result/"
)

// Inference 渲染页面
func Inference(ctx *gin.Context) {
	//// 1.接收参数
	//rgb_name := ctx.PostForm("rgb_name")
	//sar_name := ctx.PostForm("sar_name")
	//rgb_save_path := image_save_path + rgb_name
	//sar_save_path := image_save_path + sar_name
	//
	//isoverwrite := ctx.PostForm("isoverwrite")
	//if rgb_name == "" || sar_name == "" {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"status": "请上传对应的图像",
	//		"result": "",
	//	})
	//	return
	//}
	//// 2.写入数据库
	//if isoverwrite == "0" {
	//	code := model.CheckImageInfer(rgb_name, 1)
	//	//code2 := model.CheckImageUserdb(sar_name, 1)
	//	// 名称重复，返回选择是否覆盖
	//	if code == errmsg.ERROR_IMAGE_EXIT {
	//		ctx.JSON(http.StatusOK, gin.H{
	//			"status": code,
	//			"msg":    errmsg.GetErrMsg(code),
	//		})
	//		return
	//	} else {
	//		// 名称不存在，写入数据库
	//		result_name, result_path := model.GetResult(rgb_save_path, sar_save_path, rgb_name, result_save_path)
	//		result_path = strings.Split(result_path, "./")[1]
	//		_ = model.AddImage(rgb_name, base_path+image_save_path+rgb_name)
	//		_ = model.AddImage(sar_name, base_path+image_save_path+sar_name)
	//		code := model.AddInference(rgb_name, sar_name, result_name, base_path+result_path)
	//		if code != errmsg.SUCCESS {
	//			ctx.JSON(http.StatusOK, gin.H{
	//				"status": "推理失败",
	//				"result": "",
	//			})
	//			return
	//		}
	//		ctx.JSON(http.StatusOK, gin.H{
	//			"status": code,
	//			"result": base_path + result_path,
	//		})
	//		return
	//	}
	//}
	//// 该名称图像存在，用户选择覆盖，更新数据库
	//result_name, result_path := model.GetResult(rgb_save_path, sar_save_path, rgb_name, result_save_path)
	//result_path = strings.Split(result_path, "./")[1]
	//code := model.UpdateInference(rgb_name, sar_name, result_name, base_path+result_path)
	//model.UpdateImageUserdb(rgb_name, base_path+rgb_save_path)
	//model.UpdateImageUserdb(sar_name, base_path+sar_save_path)
	//if code != errmsg.SUCCESS {
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"status": "推理失败",
	//		"result": "",
	//	})
	//	return
	//}
	//ctx.JSON(http.StatusOK, gin.H{
	//	"status": code,
	//	"result": base_path + result_path,
	//})

	rgb_names := ctx.PostForm("rgb_name")
	rgb_namess := strings.Split(rgb_names, "/")
	sar_name := ctx.PostForm("sar_name")
	var rgb_save_path []string
	for _, rgb_name := range rgb_namess {
		rgb_save_path = append(rgb_save_path, image_save_path+rgb_name)
	}
	sar_save_path := image_save_path + sar_name
	isoverwrite := ctx.PostForm("isoverwrite")
	if len(rgb_namess) == 0 || sar_name == "" {
		ctx.JSON(http.StatusOK, gin.H{
			"status": "请上传对应的图像",
			"result": "",
		})
		return
	}

	// 2.写入数据库
	var code int
	var result_path []string

	if isoverwrite == "0" {
		for _, rgb_name := range rgb_namess {
			code = model.CheckImageInfer(rgb_name, 1)
			if code == errmsg.ERROR_IMAGE_EXIT {
				break
			}
		}
		//code2 := model.CheckImageUserdb(sar_name, 1)
		// 名称重复，返回选择是否覆盖
		if code == errmsg.ERROR_IMAGE_EXIT {
			ctx.JSON(http.StatusOK, gin.H{
				"status": code,
				"msg":    errmsg.GetErrMsg(code),
			})
			return
		} else {
			// 名称不存在，写入数据库
			for i := 0; i < len(rgb_namess); i++ {
				name, path := model.GetResult(rgb_save_path[i], sar_save_path, rgb_namess[i], result_save_path)
				path = strings.Split(path, "./")[1]
				result_path = append(result_path, base_path+path)
				_ = model.AddImage(rgb_namess[i], base_path+image_save_path+rgb_namess[i])
				code = model.AddInference(rgb_namess[i], sar_name, name, base_path+path)
				if code != errmsg.SUCCESS {
					ctx.JSON(http.StatusOK, gin.H{
						"status": "推理失败",
						"result": "",
					})
					return
				}
			}
			_ = model.AddImage(sar_name, base_path+image_save_path+sar_name)
			ctx.JSON(http.StatusOK, gin.H{
				"status": code,
				"result": result_path,
			})
			return
		}
	}
	// 该名称图像存在，用户选择覆盖，更新数据库
	for i := 0; i < len(rgb_namess); i++ {
		name, path := model.GetResult(rgb_save_path[i], sar_save_path, rgb_namess[i], result_save_path)
		path = strings.Split(path, "./")[1]
		code_rgb := model.CheckImageUserdb(rgb_namess[i], 1)
		code_res := model.CheckImageInfer(rgb_namess[i], 1)
		result_path = append(result_path, base_path+path)
		if code_rgb == errmsg.ERROR_IMAGE_EXIT {
			_ = model.UpdateImageUserdb(rgb_namess[i], base_path+rgb_save_path[i])
		} else {
			_ = model.AddImage(rgb_namess[i], base_path+image_save_path+rgb_namess[i])
		}
		if code_res == errmsg.ERROR_IMAGE_EXIT {
			code = model.UpdateInference(rgb_namess[i], sar_name, name, base_path+path)
		} else {
			code = model.AddInference(rgb_namess[i], sar_name, name, base_path+path)
		}
		if code != errmsg.SUCCESS {
			ctx.JSON(http.StatusOK, gin.H{
				"status": "推理失败",
				"result": "",
			})
			return
		}
	}
	code_sar := model.CheckImageUserdb(sar_name, 1)
	if code_sar == errmsg.ERROR_IMAGE_EXIT {
		_ = model.UpdateImageUserdb(sar_name, base_path+sar_name)
	} else {
		_ = model.AddImage(sar_name, base_path+image_save_path+sar_name)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": code,
		"result": result_path,
	})
}

// InferencePage 执行推理任务
func InferencePage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "inference.html", nil)
}
