package controller

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	"irisDemo/CMSProject/service"
	"irisDemo/CMSProject/utils"
)

type OrderController struct {
	Ctx iris.Context

	Service service.OrderService

	Session sessions.Session
}

/**
 * 获取所有订单数
 * /bos/orders/count
 * Get
 */
func (oc *OrderController) GetCount() mvc.Result {
	result, err := oc.Service.GetOrderCount()

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
