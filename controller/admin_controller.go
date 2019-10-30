package controller

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"irisDemo/CMSProject/service"
)

type AdminController struct {
	//Iris框架自动为每个请求添加上下文
	Ctx iris.Context

	//admin功能实体
	Service service.AdminService

	//session对象
	Session *sessions.Session

}

const (
	ADMINTABLENAME = "admin"
	ADMIN = "admin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

/**
 * 管理员登录功能
 */
func (ac *AdminController) PostLogin () mvc.Result {

	iris.New().Logger().Info("admin login")

	var adminLogin AdminLogin
	ac.Ctx.ReadJSON(&adminLogin)

	//数据校验
	if adminLogin.UserName == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": 1,
				"success": "登录失败",
				"message": "用户名或密码为空，请重新填写后尝试登录",
			},
		}
	}

	//根据用户名和密码查询对应数据库中的信息
	admin, exist := ac.Service.GetByAdminNameAndPwd(adminLogin.UserName, adminLogin.Password)
	if !exist {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": 2,
				"success": "登录失败",
				"message": "用户名或密码错误，请重新填写后尝试登录",
			},
		}
	}

	//管理员存在，设置session
	userByte, _ := json.Marshal(admin)
	ac.Session.Set(ADMIN, userByte)

	return mvc.Response{
		Object:map[string]interface{}{
			"status": 0,
			"success": "登录成功",
			"message": "管理员登录成功",
		},
	}

}