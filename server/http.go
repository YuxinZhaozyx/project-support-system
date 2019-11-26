package server

import (
	"github.com/kataras/iris/v12"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"

	server_task "project-support-system/server/task"
)

func Run(address string) {
	app := router()
	app.Run(iris.Addr(address), iris.WithoutServerError(iris.ErrServerClosed))
}

func router() *iris.Application {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything
		AllowCredentials: true,
	}))

	// 设置各模块的路由
	server_task.SetRouter(app)

	// 人员管理
	server_user.SetRouter(app)

	return app
}
