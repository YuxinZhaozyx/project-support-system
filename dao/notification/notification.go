package notification

import (
	"database/sql"
	"project-support-system/dao"

)

type NotificationDao struct {
	db *sql.DB
}

func Connect() (*NotificationDao, error) {
	var err error
	notificationDao := &NotificationDao{}
	notificationDao.db, err = dao.Connect()
	if err != nil {
		return nil, err
	}

	return notificationDao, nil
}

func (notificationDao *NotificationDao) Close() {
	notificationDao.db.Close()
}