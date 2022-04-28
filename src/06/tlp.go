package main

import (
	"html/template"
	"net/http"
)

func testTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\06\\index.html")
	t.Execute(w, "Template test")
}

func main() {
	http.HandleFunc("/template", testTemplate)
	http.ListenAndServe(":8081", nil)
}
