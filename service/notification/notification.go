package notification

import (
	model_notification "project-support-system/model/notification"
	dao_notification "project-support-system/dao/notification"
)

func CreateNotification(idUser int, title string, content string) error {
	dao, err := dao_notification.Connect()

	if err != nil{
		return err
	}
	defer dao.Close()

	err = dao.CreateNotification(idUser, title, content)

	if err!=nil{
		return err
	}else{
		return nil
	}
}

func GetNotificationsByIdByRange(idUser int, page int, num int) ([]model_notification.Notification, error){
	dao, err := dao_notification.Connect()

	if err != nil{
		return nil, err
	}
	defer dao.Close()

	var notifications []model_notification.Notification
	begin := page*num
	end := begin + num 

	notifications, err = dao.GetNotificationsByIdByRange(idUser, begin, end)
	if err!=nil{
		return nil,err
	}

	return notifications, nil
}