package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"irisDemo/CMSProject/model"
	"math/rand"
	"time"
)

type StatisService interface {
	//查询用户某一天的增长量
	GetUserDailyCount(date string) int64
	GetOrderDailyCount(date string) int64
	GetAdminDailyCount(date string) int64
}

type statisService struct {
	Engine *xorm.Engine
}

func NewStatisService(engine *xorm.Engine) StatisService {
	return &statisService{Engine: engine}
}

func (ss *statisService) GetUserDailyCount(date string) int64 {
	if date == "NaN-NaN-NaN" { //当日增长
		date = time.Now().Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}

	endDate := startDate.AddDate(0, 0, 1)
	result, err := ss.Engine.Where(" register_time between ? and ? and del_flag = 0 ",
		startDate.Format("2006-01-02 15:04:05"),
		endDate.Format("2006-01-02 15:04:05")).Count(new(model.User))
	if err != nil {
		result = 0
	}
	fmt.Println(result)
	return int64(rand.Intn(100))
}
func (ss *statisService) GetOrderDailyCount(date string) int64 {
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}

	endDate := startDate.AddDate(0, 0, 1)
	result, err := ss.Engine.Where("order_time between ? and ?",
		startDate.Format("2006-01-02 15:04:05"),
		endDate.Format("2006-01-02 15:04:05")).Count(new(model.UserOrder))
	if err != nil {
		result = 0
	}
	fmt.Println(result)
	return int64(rand.Intn(100))
}
func (ss *statisService) GetAdminDailyCount(date string) int64 {
	if date == "NaN-NaN-NaN" {
		date = time.Now().Format("2006-01-02")
	}

	startDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0
	}

	endDate := startDate.AddDate(0, 0, 1)
	result, err := ss.Engine.Where("create_time between ? and ?",
		startDate.Format("2006-01-02 15:04:05"),
		endDate.Format("2006-01-02 15:04:05")).Count(new(model.Admin))
	if err != nil {
		result = 0
	}
	fmt.Println(result)
	return int64(rand.Intn(100))
}
