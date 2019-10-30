package config

import (
	"os"
	"encoding/json"
)

//服务端配置
type AppConfig struct {
	AppName string `json:"app_name"`
	Port string `json:"port"`
	StaticPath string `json:"static_path"`
	Mode string `json:"mode"`      //环境设置，开发环境，正式环境
}

var ServConfig AppConfig

func InitConfig() *AppConfig {
	file, err := os.Open("./config.json")
	if err != nil {
		panic(err.Error())
	}
	decoder := json.NewDecoder(file)
	conf := AppConfig{}
	err = decoder.Decode(&conf)
	if err != nil {
		panic(err.Error())
	}
	return &conf
}
