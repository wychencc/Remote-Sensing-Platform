package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"version1/model"
	"version1/utils/errmsg"
)

// Login 用户登录界面
func Login(ctx *gin.Context) {
	//渲染html
	ctx.HTML(http.StatusOK, "login.html", nil)

}

// LoginCheck 登录信息校验
func LoginCheck(ctx *gin.Context) {
	//1. 接收参数
	var user model.ParamUser
	_ = ctx.ShouldBind(&user)

	//2. 业务处理
	code := model.CheckUser(user.Email)
	if code == errmsg.ERROR_USER_NOT_EXIST {
		ctx.HTML(http.StatusOK, "login.html", gin.H{
			"status": code,
			"data":   user,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	} else {
		code := model.Login(&user)
		if code == errmsg.ERROR_PASSWORD_WRONG {
			ctx.HTML(http.StatusOK, "login.html", gin.H{
				"status": code,
				"data":   user,
				"msg":    errmsg.GetErrMsg(code),
			})
			return
		}
	}
	//3. 返回响应
	ctx.Redirect(http.StatusMovedPermanently, "/api/v1/index")
}

// Register 用户注册界面
func Register(ctx *gin.Context) {
	//渲染html
	ctx.HTML(http.StatusOK, "register.html", nil)
}

// RegisterCheck 注册信息校验
func RegisterCheck(ctx *gin.Context) {
	//1. 接收参数
	var user model.User
	_ = ctx.ShouldBind(&user)

	//2. 业务处理
	code := model.CheckUser(user.Email)
	if code == errmsg.ERROR_USEREMAIL_USED {
		ctx.HTML(http.StatusOK, "register.html", gin.H{
			"status": code,
			"data":   user,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	} else {
		code := model.Register(&user)
		if code == errmsg.ERROR {
			ctx.HTML(http.StatusOK, "register.html", nil)
			return
		}
	}
	//3. 返回响应
	ctx.Redirect(http.StatusMovedPermanently, "/api/v1/login")
}

// Index 首页
func Index(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", nil)
}
