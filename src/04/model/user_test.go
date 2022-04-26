package model

import "testing"

func TestUser(t *testing.T) {
	t.Run("测试添加用户", TestAddUser)
}

func TestAddUser(t *testing.T) {
	user := &User{}
	user.AddUser()
	user.AddUser2()
}
