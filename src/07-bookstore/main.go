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
	http.Handle("/static", http.StripPrefix("static", http.FileServer(http.Dir("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\static"))))
	http.Handle("/pages", http.StripPrefix("pages", http.FileServer(http.Dir("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages"))))
	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/login", )
	http.ListenAndServe(":8081", nil)
}
