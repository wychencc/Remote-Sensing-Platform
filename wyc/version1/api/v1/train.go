package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Train 训练页面渲染
func Train(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "train.html", nil)
}
