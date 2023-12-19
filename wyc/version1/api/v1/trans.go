package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"version1_copy/model"
)

const (
	save_path  = "./static/inference/"
	trans_path = "./static/trans/"
)

func Transpose(ctx *gin.Context) {

	//file, _ := ctx.FormFile("file")
	//data_type := ctx.PostForm("data_type")
	//image_name := ctx.PostForm("image_name")
	//// 保存接收到的tif图像
	//if image_name != "" {
	//	file_name := strings.Split(image_name, ".")[0]
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"status": 200,
	//		"imgUrl": "http://localhost:6060/static/trans/" + file_name + ".png",
	//	})
	//	return
	//} else {
	//	file_name := strings.Split(file.Filename, ".")[0]
	//	_ = ctx.SaveUploadedFile(file, save_path+file.Filename)
	//	model.GetTrans(save_path+file.Filename, trans_path+file_name, data_type)
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"status": 200,
	//		"imgUrl": "http://localhost:6060/static/trans/" + file_name + ".png",
	//	})
	//	return
	//}

	files, _ := ctx.MultipartForm()
	filess := files.File["file"]
	data_type := ctx.PostForm("data_type")
	image_name := ctx.PostForm("image_name")

	var imgUrl []string
	if image_name != "" {
		//file_name := strings.Split(image_name, ".")[0]
		//ctx.JSON(http.StatusOK, gin.H{
		//	"status": 200,
		//	"imgUrl": "http://localhost:6060/static/trans/" + file_name + ".png",
		//})
		file_names := strings.Split(image_name, "/")
		for _, file_name := range file_names {
			file_name = strings.Split(file_name, ".")[0] + ".png"
			url := "http://localhost:6060/static/trans/" + file_name
			imgUrl = append(imgUrl, url)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": 200,
			"imgUrl": imgUrl,
		})
		return
	} else {
		for i := 0; i < len(filess); i++ {
			file_name := strings.Split(filess[i].Filename, ".")[0]
			_ = ctx.SaveUploadedFile(filess[i], save_path+filess[i].Filename)
			model.GetTrans(save_path+filess[i].Filename, trans_path+file_name, data_type)
			url := "http://localhost:6060/static/trans/" + file_name + ".png"
			imgUrl = append(imgUrl, url)
		}
		ctx.JSON(http.StatusOK, gin.H{
			"status": 200,
			"imgUrl": imgUrl,
		})
		return
	}
}
