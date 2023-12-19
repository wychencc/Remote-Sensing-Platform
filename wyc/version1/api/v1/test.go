package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func TestPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "test.html", nil)
}

func Test(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	fs := form.File["fs"]
	for _, f := range fs {
		log.Println(f.Filename)
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("%d uploaded!", len(fs)),
	})
}
