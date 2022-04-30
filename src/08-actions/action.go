package main

import (
	"html/template"
	"net/http"
)

func testIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\08-actions\\index.html"))
	age := 17
	t.Execute(w, age > 18)
}

func main() {
	http.HandleFunc("/testif", testIf)
	http.ListenAndServe(":8081", nil)
}
