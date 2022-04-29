package dao

import (
	"atguigu_go_web/src/07-bookstore/model"
	util "atguigu_go_web/src/07-bookstore/utils"
)

func CheckUsernameAndPassword(username, password string) (*model.User, error) {
	sqlStr := "select id,username,password,email form users where username = ? and password = ?"
	row := util.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(user.ID, user.Username, user.Password, user.Email)
	return user, nil
}

func CheckUsername(username string) (*model.User, error) {
	sqlStr := "select id,username,password,email form users where username = ?"
	row := util.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(user.ID, user.Username, user.Password, user.Email)
	return user, nil
}

func SaveUser(username, password, email string) error {
	sqlStr := "insert into users(username,password,email) values (?,?,?)"
	_, err := util.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
