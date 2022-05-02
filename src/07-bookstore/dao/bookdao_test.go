package dao

import (
	"atguigu_go_web/src/07-bookstore/model"
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	t.Run("book", TestGetBooks)
	t.Run("", TestAddBook)
	t.Run("", TestDeleteBook)
	t.Run("", TestGetBookById)
}

func TestGetBooks(t *testing.T) {
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%v个图书是%v\n", k+1, v)
	}
}

func TestAddBook(t *testing.T) {
	//title,author,price,sales,stock,img_path
	b := &model.Book{
		1,
		"test",
		"non",
		12.3,
		120,
		121,
		"",
	}
	AddBook(b)
}

func TestDeleteBook(t *testing.T) {
	DeleteBook("1")
}

func TestGetBookById(t *testing.T) {
	book, _ := GetBookById("32")
	fmt.Println(book)
}

func TestUpdateBook(t *testing.T) {
	book := &model.Book{
		32,
		"shushu",
		"luo",
		12,
		12,
		123,
		"test.png",
	}
	UpdateBook(book)
}