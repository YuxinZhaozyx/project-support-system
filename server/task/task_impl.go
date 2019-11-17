package task

import (
	"github.com/kataras/iris/v12"
)

func getTasks(ctx *iris.Context) {
	idProject := ctx.FormValue("idProject")
}

func updateTasks(ctx *iris.Context) {
	idProject := ctx.FormValue("idProject")
	taskJSON := ctx.FormValue("tasks")

}
