package controller

import (
	"atguigu_go_web/src/07-bookstore/dao"
	"html/template"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, _ := dao.GetBooks()
	t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages\\manager\\book_manager.html"))
	t.Execute(w, books)
}
