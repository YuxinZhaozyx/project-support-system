package task

import (
	"github.com/kataras/iris/v12"

	model_task "project-support-system/model/task"
	service_task "project-support-system/service/task"
	"project-support-system/utils"

	"encoding/json"
	"time"

	"strconv"
)

func getTasks(ctx iris.Context) {
	idProject, err := strconv.Atoi(ctx.FormValue("idProject"))
	tasks, err := service_task.GetTasks(idProject)
	if err != nil {
		utils.SendErrorResponse(ctx, "error in getting tasks")
		return
	}

	numTasks := len(tasks)
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
	idProject, err := strconv.Atoi(ctx.FormValue("idProject"))
	taskJSON := ctx.FormValue("tasks")

	var taskMaps []map[string]interface{}
	err = json.Unmarshal([]byte(taskJSON), taskMaps)
	if err != nil {
		utils.SendErrorResponse(ctx, "error in parsing tasks")
		return
	}


	var tasks = make([]model_task.Task, len(taskMaps))
	for i, task := range taskMaps {
		tasks[i].Id = task["id"].(int)
		tasks[i].Rank = task["rank"].(int)
		tasks[i].Title = task["title"].(string)
		tasks[i].Description = task["description"].(string)
		tasks[i].StartTime = task["start_time"].(time.Time)
		tasks[i].EndTime = task["end_time"].(time.Time)
	}

	err = service_task.UpdateTasks(idProject, tasks)
	if err != nil {
		utils.SendErrorResponse(ctx, "error in updating tasks")
		return
	}

	utils.SendSuccessResponse(ctx, nil)
}
