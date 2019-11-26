package user

import (
	"database/sql"
	model "project-support-system/model/user"
)

func (userDao *UserDao) CreateUser(user model.User) error {
	_, err := userDao.db.Exec("INSERT INTO user(`idUser`,`name`,`pwd`,`phone`) VALUES (?,?,?,?)", user.Id, user.Name, user.Pwd, user.Phone)

	return err
}

func (userDao *UserDao) GetUserById(id int) (*model.User, error){
	row := userDao.db.QueryRow("SELECT `idUser`,`name`,`pwd`,`phone` FROM `user` WHERE idUser=?", id)
	user := model.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Pwd, &user.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (userDao *UserDao) UpdateUser(user model.User) error {
	_, err := userDao.db.Exec("UPDATE `User` SET `name`=?,`pwd`=?,`phone`=?  WHERE `idUser`=?", user.Name, user.Pwd, user.Phone, user.Id)
	if err != nil {
		return err
	}
	return nil
}

