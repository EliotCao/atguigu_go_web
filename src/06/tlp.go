package main

import (
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	//t, _ := template.ParseFiles("C:\\Users\\Desktop\\atguigu_go_web\\src\\06\\index.html")
	//t.Execute(w, "Template test")

	t := template.Must(template.ParseFiles("C:\\Users\\Desktop\\atguigu_go_web\\src\\06\\index.html", "C:\\Users\\Desktop\\atguigu_go_web\\src\\06\\index2.html"))
	t.ExecuteTemplate(w, "index2.html", "hello template2")
}

func main() {
	http.HandleFunc("/template", testTemplate)
	http.ListenAndServe(":8081", nil)
}
