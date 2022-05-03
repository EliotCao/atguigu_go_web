package dao

import (
	"atguigu_go_web/src/07-bookstore/model"
	util "atguigu_go_web/src/07-bookstore/utils"
	"strconv"
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

func DeleteBook(id string) error {
	sqlStr := "delete from books where id = ?"
	_, err := util.Db.Exec(sqlStr, id)
	if err != nil{
		return err
	}
	return nil
}

func GetBookById(bookId string) (*model.Book, error) {
	sqlStr := "select id,title,author,price,sales,stock,img_path from books where id = ?"
	row := util.Db.QueryRow(sqlStr, bookId)
	book := &model.Book{}
	row.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock)
	return book, nil
}

func UpdateBook(book *model.Book) error {
	sqlStr := "update books set title=?,author=?,price=?,sales=? where id=?"
	_, err := util.Db.Exec(sqlStr, book.Title, book.Author, book.Price, book.Sales, book.Id)
	if err != nil {
		return err
	}
	return nil
}

func GetPageBooks(pageNo string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	sqlStr := "select count(*) from books"
	var totalRecord int64
	row := util.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	var pageSize int64 = 4
	var totalPageNo int64
	if totalRecord % pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	}else {
		totalPageNo = totalRecord / pageSize + 1
	}
	selStr := "select id,title,author,price,sales,stock,img_path from books limit ?,?"
	rows, _ := util.Db.Query(selStr, (iPageNo-1)*pageSize, pageSize)
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock)
		books = append(books, book)
	}
	page := &model.Page{
		books,
		iPageNo,
		pageSize,
		totalPageNo,
		totalRecord,
	}
	return page, nil
}

func GetPageBooksByPrice(pageNo, minPrice, maxPrice string) (*model.Page, error) {
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	sqlStr := "select count(*) from books where between ? and ?"
	var totalRecord int64
	row := util.Db.QueryRow(sqlStr, minPrice, maxPrice)
	row.Scan(&totalRecord)
	var pageSize int64 = 4
	var totalPageNo int64
	if totalRecord % pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	}else {
		totalPageNo = totalRecord / pageSize + 1
	}
	selStr := "select id,title,author,price,sales,stock,img_path from books between ? and ? limit ?,?"
	rows, _ := util.Db.Query(selStr,minPrice, maxPrice, (iPageNo-1)*pageSize, pageSize)
	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		rows.Scan(&book.Id, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock)
		books = append(books, book)
	}
	page := &model.Page{
		books,
		iPageNo,
		pageSize,
		totalPageNo,
		totalRecord,
	}
	return page, nil
}