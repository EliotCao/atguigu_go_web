package dao

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	t.Run("test 1", TestLogin)
	t.Run("test 2", TestCheckUsername)
	t.Run("test 3", TestSaveUser)

}

func TestLogin(t *testing.T) {
	user, _ := CheckUsernameAndPassword("admin", "123456")
	fmt.Println("", user)
}

func TestCheckUsername(t *testing.T) {
	user, _ := CheckUsername("admin")
	fmt.Println("", user)
}

func TestSaveUser(t *testing.T) {
	err := SaveUser("admin", "123456", "admin@go.com")
	fmt.Println("", err)
}
