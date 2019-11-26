package notification

import (
	"github.com/kataras/iris/v12"
	"project-support-system/utils"
	service_notification "project-support-system/service/notification"
	"strconv"
)

func getNotification(ctx iris.Context) {
	idUser, err := strconv.Atoi(ctx.FormValue("idUser"))
	if err!=nil{
		utils.SendErrorResponse(ctx, "idUser error")
		return
	}

	page, err := strconv.Atoi(ctx.FormValue("page"))
	if err!=nil{
		utils.SendErrorResponse(ctx, "page number error")
		return
	}

	count, err := strconv.Atoi(ctx.FormValue("count"))
	if err!=nil{
		utils.SendErrorResponse(ctx, "count number error")
		return
	}

	notifications, err := service_notification.GetNotificationsByIdByRange(idUser, page, count)
	if err!=nil{
		utils.SendErrorResponse(ctx, "can't get notifications")
		return
	}else{
		data := make([]map[string]interface{}, count)
		
		for i:=0; i<count; i++{
			data[i] = map[string]interface{}{
				"num": i,
				"title": notifications[i].Title,
				"content": notifications[i].Content,
			}
		}
		utils.SendSuccessResponse(ctx, data)
		return
	}
}

func createNotification(ctx iris.Context) {
	idUser, err := strconv.Atoi(ctx.FormValue("idUser"))
	if err!=nil{
		utils.SendErrorResponse(ctx, "idUser error")
		return
	}

	title  := ctx.FormValue("title")
	content := ctx.FormValue("content")

	err = service_notification.CreateNotification(idUser, title, content);

	if err!=nil{
		utils.SendErrorResponse(ctx, "can't create now")
		return
	}else{
		utils.SendSuccessResponse(ctx, nil)
		return
	}
}
