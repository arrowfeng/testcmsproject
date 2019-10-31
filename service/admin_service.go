package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"irisDemo/CMSProject/model"
	"math/rand"
)

type AdminService interface {
	//校验管理员信息
	GetByAdminNameAndPwd(string, string) (model.Admin, bool)
	//获取管理员数量
	GetAdminCount() (int64, error)
}

type adminService struct {
	engine *xorm.Engine
}

func NewAdminService(engine *xorm.Engine) AdminService {
	return &adminService{engine: engine}
}

func (as *adminService) GetByAdminNameAndPwd(name, pwd string) (model.Admin, bool) {
	var admin model.Admin
	as.engine.Where("admin_name = ?", name).And("pwd = ?", pwd).Get(&admin)
	return admin, admin.AdminId != 0
}

func (as *adminService) GetAdminCount() (int64, error) {
	result, err := as.engine.Count(new(model.Admin))

	if err != nil {
		panic(err.Error())
		return 0, err
	}

	fmt.Println(result)
	return int64(rand.Intn(100)), nil
}
