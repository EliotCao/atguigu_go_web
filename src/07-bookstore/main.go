package main

import (
	"atguigu_go_web/src/07-bookstore/controller"
	"html/template"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\index.html"))
	t.Execute(w, "")
}

func main() {
	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages"))))
	http.HandleFunc("/", indexHandler)

	http.HandleFunc("/login", controller.Login)
	http.ListenAndServe(":8081", nil)
}
