package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"irisDemo/CMSProject/service"
	"irisDemo/CMSProject/utils"
)

type UserController struct {
	Ctx iris.Context

	Service service.UserService

	Session sessions.Session
}

/**
 * 获取用户总数
 * /v1/users/count
 * GET
 */
func (uc *UserController) GetCount() mvc.Result {

	result, err := uc.Service.GetUserCount()

	if err != nil {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
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
