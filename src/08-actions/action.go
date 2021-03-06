package main

import (
	"atguigu_go_web/src/08-actions/model"
	"html/template"
	"net/http"
)

func testIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("C:\\Users\\Desktop\\atguigu_go_web\\src\\08-actions\\if.html"))
	age := 17
	t.Execute(w, age > 18)
}

func testRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("C:\\Users\\Desktop\\atguigu_go_web\\src\\08-actions\\range.html"))
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

func testWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("C:\\Users\\Desktop\\atguigu_go_web\\src\\08-actions\\with.html"))
	t.Execute(w, "狸猫")
}

func testTemplate(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("C:\\Users\\Desktop\\atguigu_go_web\\src\\08-actions\\template1.html", "C:\\Users\\Desktop\\atguigu_go_web\\src\\08-actions\\template2.html"))
	t.Execute(w, "Template data test")
}

func testDefine(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("C:\\Users\\Desktop\\atguigu_go_web\\src\\08-actions\\define.html"))
	t.ExecuteTemplate(w, "model", "")
}

func testDefine2(w http.ResponseWriter, r *http.Request) {
	age := 19
	var t *template.Template
	if age > 18 {
		t = template.Must(template.ParseFiles("C:\\Users\\Desktop\\atguigu_go_web\\src\\08-actions\\define2.html"))
	}else {
		t = template.Must(template.ParseFiles("C:\\Users\\Desktop\\atguigu_go_web\\src\\08-actions\\define2.html", "C:\\Users\\Desktop\\atguigu_go_web\\src\\08-actions\\content2.html"))
	}
	t.ExecuteTemplate(w, "model", "")
}

func main() {
	http.HandleFunc("/testif", testIf)
	http.HandleFunc("/testrange", testRange)
	http.HandleFunc("/testwith", testWith)
	http.HandleFunc("/testtemplate", testTemplate)
	http.HandleFunc("/testdefine", testDefine)
	http.HandleFunc("/testdefine2", testDefine2)

	http.ListenAndServe(":8081", nil)
}
