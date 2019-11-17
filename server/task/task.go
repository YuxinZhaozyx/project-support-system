package task

import (
	"github.com/kataras/iris/v12"
)

func SetRouter(app *iris.Application) {
	taskParty := app.Party("/task")
	{
		taskParty.Post("/getall", getTasks)
		taskParty.Post("/update", updateTasks)
	}
}
