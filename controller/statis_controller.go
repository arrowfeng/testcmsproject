package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"irisDemo/CMSProject/service"
	"irisDemo/CMSProject/utils"
	"strings"
)

type StatisController struct {
	Ctx iris.Context

	Service service.StatisService

	Session *sessions.Session
}

var (
	ADMINMODULE = "ADMIN_"
	USERMODULE  = "USER_"
	ORDERMODULE = "ORDER_"
)

/**
 *获取各模块的增长数量
 *Get
 *path: /statis/{model}/{date}/count
 */
func (sc *StatisController) GetCount() mvc.Result {

	path := sc.Ctx.Path()
	var pathSlice []string
	if path != "" {
		pathSlice = strings.Split(path, "/")
	}
	//  /statis/{model}/{date}/count
	//  "" "statis" "{model}" "{date}" "count"
	//不符合请求格式
	if len(pathSlice) != 5 {
		return mvc.Response{
			Object: map[string]interface{}{
				"status": utils.RECODE_FAIL,
				"count":  0,
			},
		}
	}

	//解析请求
	pathSlice = pathSlice[1:]
	model := pathSlice[1]
	date := pathSlice[2]
	var result int64
	switch model {
	case "admin":
		adminResult := sc.Session.Get(ADMINMODULE + date)
		if adminResult != nil {
			adminResult = adminResult.(int64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  adminResult,
				},
			}
		} else {
			iris.New().Logger().Error(date) //没有找到缓存
			result = sc.Service.GetAdminDailyCount(date)
			//设置缓存
			sc.Session.Set(ADMINMODULE+date, result)
		}
	case "user":
		userResult := sc.Session.Get(USERMODULE + date)
		if userResult != nil {
			userResult = userResult.(int64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  userResult,
				},
			}
		} else {
			iris.New().Logger().Error(date) //没有找到缓存
			result = sc.Service.GetUserDailyCount(date)
			//设置缓存
			sc.Session.Set(USERMODULE+date, result)
		}
	case "order":
		orderResult := sc.Session.Get(ORDERMODULE + date)
		if orderResult != nil {
			orderResult = orderResult.(int64)
			return mvc.Response{
				Object: map[string]interface{}{
					"status": utils.RECODE_OK,
					"count":  orderResult,
				},
			}
		} else {
			iris.New().Logger().Error(date) //没有找到缓存
			result = sc.Service.GetOrderDailyCount(date)
			//设置缓存
			sc.Session.Set(ORDERMODULE+date, result)
		}
	default:
		iris.New().Logger().Error("请求模块错误")
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
