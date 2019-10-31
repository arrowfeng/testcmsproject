package controller

import (
	"encoding/json"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"irisDemo/CMSProject/model"
	"irisDemo/CMSProject/service"
	"irisDemo/CMSProject/utils"
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
	ADMIN          = "admin"
)

type AdminLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

/**
 * 管理员登录功能
 * /admin/login
 * post
 */
func (ac *AdminController) PostLogin() mvc.Result {

	iris.New().Logger().Info("admin login")

	var adminLogin AdminLogin
	ac.Ctx.ReadJSON(&adminLogin)

	//数据校验
	if adminLogin.UserName == "" || adminLogin.Password == "" {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
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
				"status":  utils.RECODE_FAIL,
				"success": "登录失败",
				"message": "用户名或密码错误，请重新填写后尝试登录",
			},
		}
	}

	//管理员存在，设置session
	userByte, _ := json.Marshal(admin)
	ac.Session.Set(ADMIN, userByte)

	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"success": "登录成功",
			"message": "管理员登录成功",
		},
	}

}

/**
 * 管理员退出功能
 * /admin/singout
 * get
 */
func (ac *AdminController) GetSingout() mvc.Result {
	ac.Session.Delete(ADMIN)
	return mvc.Response{
		Object: map[string]interface{}{
			"status":  utils.RECODE_OK,
			"success": utils.Recode2Text(utils.RESPMSG_SIGNOUT),
		},
	}
}

/**
 *获取管理员信息接口
 *请求类型：GET
 *请求url：/admin/info
 */
func (ac *AdminController) GetInfo() mvc.Result {

	//从session获取管理员信息
	userByte := ac.Session.Get(ADMIN)

	//session为空
	if userByte == nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.ERROR_UNLOGIN,
				"message": utils.Recode2Text(utils.ERROR_UNLOGIN),
			},
		}
	}

	//解析数据到admin数据结构
	var admin model.Admin
	err := json.Unmarshal(userByte.([]byte), &admin)
	//解析失败
	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_UNLOGIN,
				"type":    utils.ERROR_UNLOGIN,
				"message": utils.Recode2Text(utils.ERROR_UNLOGIN),
			},
		}
	}

	//解析成功
	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"data":   admin.AdminToRespDesc(),
		},
	}
}

/**
 * 获取管理员总数
 * 请求类型：GET
 * 请求url：/admin/count
 */
func (ac *AdminController) GetCount() mvc.Result {

	result, err := ac.Service.GetAdminCount()

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status":  utils.RECODE_FAIL,
				"message": utils.Recode2Text(utils.RESPMSG_ERRORADMINCOUNT),
				"count":   0,
			},
		}
	}

	return mvc.Response{
		Object: map[string]interface{}{
			"status": utils.RECODE_OK,
			"count":  result,
		},
	}
}
