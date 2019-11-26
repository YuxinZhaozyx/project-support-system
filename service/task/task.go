package task

import (
	dao_task "project-support-system/dao/task"
	model_task "project-support-system/model/task"
)

func GetTasks(idProject int) ([]model_task.Task, error) {
	dao, err := dao_task.Connect()
	if err != nil {
		return nil, err
	}
	defer dao.Close()

	var tasks []model_task.Task
	tasks, err = dao.GetTasksByIdProject(idProject)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func UpdateTasks(tasks []model_task.Task) error {
	dao, err := dao_task.Connect()
	if err != nil {
		return nil, err
	}
	defer dao.Close()

	var task []model_task.Task
	tasksOld, err = dao.DeleteTasksByIdProject(idProject)
	if err != nil {
		return err
	}

	err = dao.CreateTasks(task)
	if err != nil {
		return err
	}

	return nil
}
