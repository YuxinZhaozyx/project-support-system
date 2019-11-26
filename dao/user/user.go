package user

import (
	"database/sql"
	dao "project-support-system/dao"
)

type UserDao struct {
	db *sql.DB
}

func Connect() (*UserDao, error) {
	var err error
	userDao := &UserDao{}
	userDao.db, err = dao.Connect()
	if err != nil {
		return nil, err
	}

	return userDao, nil
}

func (userDao *UserDao) Close() {
	userDao.db.Close()
}