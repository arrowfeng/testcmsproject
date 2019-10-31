package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"irisDemo/CMSProject/model"
	"math/rand"
)

type UserService interface {
	GetUserCount() (int64, error)
}

type userService struct {
	Engine *xorm.Engine
}

func NewUserService(engine *xorm.Engine) UserService {
	return &userService{Engine: engine}
}

func (us *userService) GetUserCount() (int64, error) {

	result, err := us.Engine.Count(new(model.User))

	if err != nil {
		panic(err.Error())
		return 0, err
	}

	fmt.Println(result)
	return int64(rand.Intn(100)), nil
}
