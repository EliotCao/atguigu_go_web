package model

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("测试添加用户", TestAddUser)
}

func TestAddUser(t *testing.T) {
	user := &User{}
	user.AddUser()
	user.AddUser2()
}

func TestGetUserById(t *testing.T) {
	user := User{
		ID: 1,
	}
	u,_ := user.GetUserById()
	fmt.Println("查询到用户信息", u)
}
