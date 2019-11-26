package notification

import (
	"time"
	model "project-support-system/model/notification"
)

func (notificationDao *NotificationDao) CreateNotification(idUser int, title string, content string) error {
	stmt, err := notificationDao.db.Prepare("INSERT INTO Notices(idNotice) VALUES(?, ?, ?, ?)")
	if err != nil{
		return err
	}

	sendTime := time.Now().Format("2006-01-02 15:04:05")
	_, err = stmt.Exec(title, content, sendTime, idUser)

	if err!=nil{
		return err
	}

	var idNotice int
	row := notificationDao.db.QueryRow("SELECT idNotice FROM Notice WHERE idSender = ? AND time = ?", idUser, sendTime)
	row.Scan(&idUser, &idNotice)

	stmt, err = notificationDao.db.Prepare("INSERT INTO User_has_Notice(User_idUser, Notice_idNotice) VALUES(?, ?)")
	if err != nil{
		return err
	}
	_, err = stmt.Exec(idUser, idNotice)

	if err!=nil{
		return err
	}

	return nil
}

func (notificationDao *NotificationDao) GetNotificationsByIdByRange(id int, begin int, end int) (notificationsRes []model.Notification, err error) {
	rows, err := notificationDao.db.Query("SELECT * FROM Notice")
	if err!=nil{
		return nil, err
	}

	var notifications []model.Notification
	var i int = 0
	for rows.Next(){
		var id int
		var title string
		var content string 
		var time time.Time
		var idSender int

		err = rows.Scan(&id, &title, &content, &time, &idSender)

		if err!=nil{
			return nil, err
		}

		notifications[i] = *model.NewNotification(id, title, content, time, idSender)
		i = i+1
	}

	return notifications, nil
}
