package utils

//请求状态码
const (
	RECODE_OK      = 1  //请求成功 正常
	RECODE_FAIL    = 0  //请求失败
	RECODE_UNLOGIN = -1 //未登录 没有权限
)

//业务逻辑状态码
const (
	RESPMSG_OK   = "1"
	RESPMSG_FAIL = "0"

	//管理员
	RESPMSG_SUCCESSLOGIN    = "SUCCESS_LOGIN"
	RESPMSG_FAILURELOGIN    = "FAILURE_LOGIN"
	RESPMSG_SUCCESSSESSION  = "SUCCESS_SESSION"
	RESPMSG_ERRORSESSION    = "FAILURE_SESSION"
	RESPMSG_SIGNOUT         = "SUCCESS_SIGNOUT"
	RESPMSG_HASNOACCESS     = "HAS_NO_ACCESS"
	RESPMSG_ERRORADMINCOUNT = "ERROR_ADMINCOUNT"

	//未登陆
	ERROR_UNLOGIN = "ERROR_UNLOGIN"

	//其他错误
	RECODE_UNKNOWER = "8000"
)

//业务逻辑描述
var recodeText = map[string]string{
	RESPMSG_OK:    "成功",
	RESPMSG_FAIL:  "失败",
	ERROR_UNLOGIN: "未登陆无操作权限，请先登陆",

	//管理员
	RESPMSG_SUCCESSLOGIN:    "管理员登录成功",
	RESPMSG_FAILURELOGIN:    "管理员账号或密码错误，登录失败",
	RESPMSG_SUCCESSSESSION:  "获取管理员信息成功",
	RESPMSG_ERRORSESSION:    "获取管理员信息失败",
	RESPMSG_SIGNOUT:         "退出成功",
	RESPMSG_HASNOACCESS:     "亲，您的权限不足",
	RESPMSG_ERRORADMINCOUNT: "获取管理员总数失败",

	//其他错误
	RECODE_UNKNOWER: "服务器未知错误",
}

func Recode2Text(respCode string) string {
	str, ok := recodeText[respCode]
	if ok {
		return str
	}
	return recodeText[RECODE_UNKNOWER]
}
