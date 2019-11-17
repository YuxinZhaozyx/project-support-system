package utils

import (
	"github.com/kataras/iris/v12"
)

func SendErrorResponse(ctx *iris.Context, errorMessage string) {
	ctx.JSON(iris.Map{
		"success": false,
		"error": errorMessage
	})
}

func SendSuccessResponse(ctx *iris.Context, data: interface{}) {
	ctx.JSON(iris.Map{
		"success": true,
		"data": data
	})
}