package dao

import (
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	t.Run("book", TestGetBooks)
}

func TestGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%v个图书是%v\n", k+1, v)
	}
}

