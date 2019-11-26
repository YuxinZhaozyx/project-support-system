package user

import (
	"github.com/kataras/iris/v12"
)

func SetRouter(app *iris.Application) {
	userParty := app.Party("/user")
	{
		userParty.Post("/create/user/request", updateUser)
		userParty.Post("/login", login)
	}
}