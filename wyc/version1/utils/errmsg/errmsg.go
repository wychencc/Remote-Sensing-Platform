package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	//code=1000.. 用户模块错误
	ERROR_USEREMAIL_USED   = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_NOT_EXIST  = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NOTADMIN    = 1008

	//code=2000.. 模型模块错误
	ERROR_MODEL_NOT_EXIST = 2001
	ERROR_MODEL_EXISTED   = 2002

	//code=3000.. 用户数据库错误
	ERROR_IMAGE_EXIT = 3001
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USEREMAIL_USED:   "邮箱以注册!",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_NOT_EXIST:  "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WRONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误",
	ERROR_USER_NOTADMIN:    "用户无权限",
	ERROR_MODEL_NOT_EXIST:  "模型不存在",
	ERROR_MODEL_EXISTED:    "模型已存在",
	ERROR_IMAGE_EXIT:       "存在相同名称图片，请返回采集",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
