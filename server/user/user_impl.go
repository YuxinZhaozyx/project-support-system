package user

import (
	"github.com/kataras/iris/v12"
	"strconv"
	"regexp"
	model_user "project-support-system/model/user"
	service_user "project-support-system/service/user"
)

func updateUser(ctx iris.Context){
	id := ctx.FormValue("id")
	name := ctx.FormValue("name")
	pwd := ctx.FormValue("pwd")
	phone := ctx.FormValue("phone")

	// 返回结果
	var success bool = false;

	// id 类型转换为idInt
	idInt, err := strconv.Atoi(id)
	
	if err != nil {
		ctx.JSON(iris.Map{
			"success":success,
			"error":"User id should be a int value",
		})
		return;
	}

	// 检查手机号是否合理
	regPhone := regexp.MustCompile(`^1[3|4|5|7|8][0-9]{9}$`) //匹配手机号码

	if !regPhone.MatchString(phone) {
		// 如果不匹配
		ctx.JSON(iris.Map{
			"success":success,
			"error":"Wrong phone number format",
		})
		return
	}

	// 密码长度需要大于8位小于16位, 判断密码合理性
	if len(pwd) < 8 || len(pwd) > 16 {
		ctx.JSON(iris.Map{
			"success":success,
			"error":"The length of password should larger than eight and smaller than sixteen",
		})
		return
	}

	// 判断名字是否为空
	if len(name) == 0 {
		ctx.JSON(iris.Map{
			"success":success,
			"error":"The name is not allowed to be a null",
		})
	}

	// 数据合理，生成User对象进行添加或更新
	user, _ := model_user.NewUser(idInt, name, pwd, phone)

	if user == nil {
		ctx.JSON(iris.Map{
			"success":success,
			"error":"Create user failed",
		})
		return
	}

	err = service_user.UpdateUser(*user)

	if err == nil {
		success = true
	}

	ctx.JSON(iris.Map{
		"success":success,
		"error":err,
	})
}

func login(ctx iris.Context) {
	id := ctx.FormValue("id")
	pwd := ctx.FormValue("pwd")

	// 返回信息
	var success bool = false

	// 转换类型
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(iris.Map{
			"success":success,
			"error":err,
		})
	}

	err = service_user.Login(idInt, pwd)
	if err == nil {
		success = true
	}

	ctx.JSON(iris.Map{
		"success":success,
		"error":err,
	})
}