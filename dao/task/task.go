package task

import (
	"database/sql"

	"project-support-system/dao"
)

type TaskDao struct {
	db *sql.DB
}

func Connect() (*TaskDao, error) {
	var err error
	taskDao := new(TaskDao)
	taskDao.db, err = dao.Connect()
	if err != nil {
		return nil, err
	}

	return taskDao, nil
}

func (taskDao *TaskDao) Close() {
	taskDao.db.Close()
}
