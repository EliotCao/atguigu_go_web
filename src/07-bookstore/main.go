package main

import (
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\index.html"))
	t.Execute(w, "")
}

func main() {
	http.HandleFunc("/index", indexHandler)
	http.ListenAndServe(":8081", nil)
}
