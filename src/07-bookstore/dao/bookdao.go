package dao

import (
	"atguigu_go_web/src/07-bookstore/model"
	util "atguigu_go_web/src/07-bookstore/utils"
)

func GetBooks() ([]*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books"
	rows, err := util.Db.Query(sqlStr)
	if err != nil {
		return nil,err
	}
	var books []*model.Book
	for rows.Next() {
		var book *model.Book
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImagePath)
		books = append(books, book)
	}
	return books, nil
}

func AddBook(b *model.Book) error {
	sqlStr := "insert into books(title,author,price,sales,stock,img_path) values(?,?,?,?,?,?)"
	_, err := util.Db.Exec(sqlStr, b.Title, b.Author, b.Price, b.Sales, b.Stock, b.ImagePath)
	if err != nil {
		return err
	}
	return nil
}
