package main

import (
	"atguigu_go_web/src/08-actions/model"
	"html/template"
	"net/http"
)

func testIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\08-actions\\index.html"))
	age := 17
	t.Execute(w, age > 18)
}

func testRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\08-actions\\range.html"))
	emps := []*model.Employee{
		{
			1,
			"Xiaolu LI",
			"xialu.li@go.com",
		},
		{
			2,
			"Sicong Wang",
			"sicong.wang@go.com",
		},
		{
			3,
			"Andy liu",
			"andy.liu@hk.com",
		},
	}
	t.Execute(w, emps)
}

func main() {
	http.HandleFunc("/testif", testIf)
	http.HandleFunc("/testrange", testRange)
	http.ListenAndServe(":8081", nil)
}
