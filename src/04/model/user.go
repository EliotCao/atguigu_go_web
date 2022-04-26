package model

import (
	"atguigu_go_web/src/04/util"
	"fmt"
)

type User struct {
	ID       int
	username string
	password string
	email    string
}

func (u User) AddUser() error {
	sqlStr := "insert into users(username, password, email) values(?,?,?)"
	inSmt, err := util.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("prepare错误", err)
	}
	_, err2 := inSmt.Exec("admin", "123456", "123456@qq.com")
	if err2 != nil {
		fmt.Println("prepare错误", err2)
		return err2
	}
	return nil
}

func (u User) AddUser2() error {
	sqlStr := "insert into users(username, password, email) values(?,?,?)"
	_, err := util.Db.Exec(sqlStr,"admin", "123456", "123456@qq.com")
	if err != nil {
		fmt.Println("prepare错误", err)
	}
	return nil
}

func (u User) GetUserById() (*User, error) {
	sqlStr := "select id, username, password, email from users where id = ?"
	row := util.Db.QueryRow(sqlStr, u.ID)
	var id int
	var username string
	var password string
	var email string
	err := row.Scan(&id, &username, &password, &email)
	if err != nil {
		return nil, err
	}
	return &User{
		id,
		username,
		password,
		email,
	}, nil
}