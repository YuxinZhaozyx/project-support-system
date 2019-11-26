package user

import (
	"errors"
	model_user "project-support-system/model/user"
	dao_user "project-support-system/dao/user"
)

func UpdateUser(user model_user.User) error {
	dao, _ := dao_user.Connect()

	if dao == nil {
		return errors.New("database goes wrong")
	}

	userUpdate, err := dao.GetUserById(user.Id)

	// 如果用户已存在
	if userUpdate != nil {
		err = dao.UpdateUser(user)
		if err != nil {
			return err
		}
	} else if err == nil {
		// 此时该用户不存在,新建
		err = dao.CreateUser(user)
		if err != nil {
			// 创建失败
			return err
		}
	} else {
		// 其他错误
		dao.Close()
		return err
	}

	dao.Close()
	return nil
}

func Login(id int, pwd string) error {
	dao, _ := dao_user.Connect()

	user, _ := dao.GetUserById(id)

	// 如果用户不存存在
	if user == nil {
		return errors.New("The user does not exist")
	}

	if pwd == user.Pwd {
		dao.Close()
		return nil
	}
	// 密码错误
	dao.Close()
	return errors.New("The password is wrong")
}