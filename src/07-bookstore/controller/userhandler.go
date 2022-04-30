package controller

import (
	"atguigu_go_web/src/07-bookstore/dao"
	"html/template"
	"net/http"
)

//Login handle user login
func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ := dao.CheckUsernameAndPassword(username, password)
	if user != nil {
		t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages\\user\\login.html"))
		t.Execute(w, "")
	}else {

	}
}
