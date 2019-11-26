package task

import (
	model "project-support-system/model/task"
)

func (taskDao *TaskDao) CreateTasks(tasks []model.Task) (err error) {
	insertString := "INSERT INTO Task (title, startTime, endTime, description, rank, prevTask, idGroup) VALUE (?, ?, ?, ?, ?, ?, ?)"

	for _, task := range tasks {
		_, err = taskDao.db.Exec(insertString, task.Title, task.StartTime, task.EndTime, task.Description, task.Rank, task.PrevTask, task.IdGroup)
		if err != nil {
			return err
		}
	}

	return nil
}

func (taskDao *TaskDao) GetTasksByIdProject(idProject int) ([]model.Task, error) {
	queryString := `SELECT idTask, title, startTime, endTime, description, rank, prevTask, idGroup FROM Task WHERE idProject=?`

	rows, err := taskDao.db.Query(queryString)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]model.Task, 0)

	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.Id, &task.Title, &task.StartTime, &task.EndTime, &task.Description, &task.Rank, &task.PrevTask, &task.IdGroup)
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (taskDao *TaskDao) DeleteTasksByIdProject(idProject int) error {
	deleteString := `DELETE FROM Task WHERE idProject = ?`

	_, err = taskDao.db.Exec(deleteString, idProject)
	if err != nil {
		return err
	}

	return nil
}
