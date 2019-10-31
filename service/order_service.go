package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"irisDemo/CMSProject/model"
	"math/rand"
)

type OrderService interface {
	GetOrderCount() (int64, error)
}

type orderService struct {
	Engine *xorm.Engine
}

func NewOrderService(engine *xorm.Engine) OrderService {
	return &orderService{Engine: engine}
}

func (os *orderService) GetOrderCount() (int64, error) {

	result, err := os.Engine.Count(new(model.UserOrder))

	if err != nil {
		panic(err.Error())
		return 0, err
	}

	fmt.Println(result)
	return int64(rand.Intn(100)), nil

}
