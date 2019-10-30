package main

import (
	"github.com/kataras/iris"
	_ "github.com/kataras/iris/_examples/mvc/login/datasource"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	_ "github.com/kataras/iris/mvc"
	"github.com/kataras/iris/sessions"
	_ "github.com/kataras/iris/sessions"
	"irisDemo/CMSProject/config"
	"irisDemo/CMSProject/controller"
	"irisDemo/CMSProject/service"
	"time"
	_ "time"
)

func main() {
	app := newApp()
	//配置应用app
	configuration(app)

	//路由设置
	mvcHandle(app)

	config := config.InitConfig()
	addr := ":" + config.Port
	app.Run(
		iris.Addr(addr),      //监听地址
		iris.WithoutServerError(iris.ErrServerClosed),     //无服务的错误提示
		iris.WithOptimizations,       //对json序列化更快
	)

}

//构建App
func newApp() *iris.Application {

	app := iris.New()

	//设置日志级别
	app.Logger().SetLevel("debug")

	//注册静态资源
	app.HandleDir("/static", "./static")
	app.HandleDir("/manage/static", "./static")
	app.HandleDir("/img", "./static/img")

	//注册视图文件
	app.RegisterView(iris.HTML("./static", ".html"))
	app.Get("/", func(context context.Context) {
		context.View("index.html")
	})

	return app
}

//应用设置
func configuration(app *iris.Application) {

	//设置编码
	app.Configure(iris.WithConfiguration(iris.Configuration{
		Charset: "UTF-8",
	}))

	//错误配置
	//未发现错误
	app.OnErrorCode(iris.StatusNotFound, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusNotFound,
			"msg":    "not found",
			"data":   iris.Map{},
		})
	})

	app.OnErrorCode(iris.StatusInternalServerError, func(context context.Context) {
		context.JSON(iris.Map{
			"errmsg": iris.StatusInternalServerError,
			"msg":    "internel error",
			"data":   iris.Map{},
		})
	})
}

//路由设置
func mvcHandle(app *iris.Application) {
	//启用session
	sessManager := sessions.New(sessions.Config{
		Cookie:  "sessioncookie",
		Expires: 24 * time.Hour,
	})

	//管理员功能模块
	adminService := service.NewAdminService()

	admin := mvc.New(app.Party("/admin"))
	admin.Register(
		adminService,
		sessManager.Start,
	)
	admin.Handle(new(controller.AdminController))

	//用户功能模块

}
