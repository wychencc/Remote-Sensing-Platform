package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"version1_copy/model"
)

func UserdbPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "userdb.html", nil)
}

func GetUserDbData(ctx *gin.Context) {

	// 1.拿到userdb里的图片名称和保存路径
	gee_names, gee_paths, _ := model.GetGeeList()

	image_names, image_paths, _ := model.GetImageList()
	res_names, res_paths, _ := model.GetResultList()

	for i := 0; i < len(image_names); i++ {
		image_paths[i] = base_path + image_tans_path + strings.Split(image_names[i], ".")[0] + ".png"
	}

	ctx.JSON(http.StatusOK, gin.H{
		"gee_names":   gee_names,
		"gee_paths":   gee_paths,
		"image_names": image_names,
		"image_paths": image_paths,
		"res_names":   res_names,
		"res_paths":   res_paths,
	})
}
