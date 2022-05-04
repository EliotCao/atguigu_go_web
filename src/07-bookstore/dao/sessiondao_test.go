package dao

import (
	"atguigu_go_web/src/07-bookstore/model"
	"fmt"
	"testing"
)

func TestSession(t *testing.T) {
	t.Run("AddSession", TestAddSession)
	t.Run("DeleteSession", TestDeleteSession)
	t.Run("GetSession", TestGetSession)
}

func TestAddSession(t *testing.T) {
	session := &model.Session{
		"test",
		"test",
		33,
	}
	AddSession(session)
}

func TestDeleteSession(t *testing.T) {
	sessionId := "test"
	DeleteSession(sessionId)
}

func TestGetSession(t *testing.T) {
	session, _ := GetSession("test")
	fmt.Println(session)
}