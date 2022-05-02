package controller

import (
	"atguigu_go_web/src/07-bookstore/dao"
	"atguigu_go_web/src/07-bookstore/model"
	"html/template"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()
	t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages\\manager\\book_manager.html"))
	t.Execute(w, books)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	price := r.PostFormValue("price")
	author := r.PostFormValue("author")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")
	fPrice, _ := strconv.ParseFloat(price, 64)
	iSales, _ := strconv.ParseInt(sales, 10,0)
	iStock, _ := strconv.ParseInt(stock, 10,0)
	b := &model.Book{
		1,
		title,
		author,
		fPrice,
		int(iSales),
		int(iStock),
		"test.png",
	}
	dao.AddBook(b)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("bookId")
	dao.DeleteBook(id)
	GetBooks(w,r)
}

func ToUpdateBookPage(w http.ResponseWriter, r *http.Request) {
	bookId := r.FormValue("bookId")
	book, _ := dao.GetBookById(bookId)
	t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages\\manager\\book_modify.html"))
	t.Execute(w, book)
}