package notification

import "time"

type Notification struct {
	Id           int       `json:"idNotification"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Time         time.Time `json:"time"`
	IdSender     int       `json:"idSender"`
}

func NewNotification(id int, title string, content string, time time.Time, idSender int) *Notification{
	notification := new(Notification)

	notification.Id = id
	notification.Title = title
	notification.Content = content
	notification.Time = time
	notification.IdSender = idSender

	return notification
}

