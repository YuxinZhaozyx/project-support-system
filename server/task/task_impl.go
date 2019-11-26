package task

import (
	"github.com/kataras/iris/v12"

	model_task "project-support-system/model/task"
	service_task "project-support-system/service/task"
	"project-support-system/utils"

	"encoding/json"
)

func getTasks(ctx iris.Context) {
	idProject := ctx.FormValue("idProject")
	tasks, err := service_task.GetTasks(idProject)
	if err != nil {
		utils.SendErrorResponse(ctx, "error in getting tasks")
		return
	}

	numTasks = len(tasks)
	data := make([]map[string]interface{}, numTasks)
	for i, task := range tasks {
		data[i] = map[string]interface{}{
			"id":          task.Id,
			"rank":        task.Rank,
			"title":       task.Title,
			"description": task.Description,
			"start_time":  task.StartTime,
			"end_time":    task.EndTime,
		}
	}

	utils.SendSuccessResponse(ctx, data)

}

func updateTasks(ctx iris.Context) {
	idProject := ctx.FormValue("idProject")
	taskJSON := ctx.FormValue("tasks")

	var taskMaps []map[string]interface{}
	err := json.Unmarshal(tasksJSON.([]byte), tasksMap)
	if err != nil {
		utils.SendErrorResponse(ctx, "error in parsing tasks")
		return
	}

	var tasks = make([]model_task.Task, len(taskMaps))
	for i, task := range taskMaps {
		tasks[i].Id = task["id"]
		tasks[i].Rank = task["rank"]
		tasks[i].Title = task["title"]
		tasks[i].Description = task["description"]
		tasks[i].StartTime = task["start_time"]
		tasks[i].EndTime = task["end_time"]
	}

	err := model_task.updateTasks(idProject, tasks)
	if err != nil {
		utils.SendErrorResponse(ctx, "error in updating tasks")
		return
	}

	utils.SendSuccessResponse(ctx, nil)
}
