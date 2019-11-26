package notification

import (
	"github.com/kataras/iris/v12"
)

func SetRouter(app *iris.Application) {
	notificationParty := app.Party("/notice")
	{
		notificationParty.Post("/get", getNotification)
		notificationParty.Post("/create", createNotification)
	}
}