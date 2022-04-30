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
	if user.ID > 0 {
		t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages\\user\\login_success.html"))
		t.Execute(w, "")
	}else {
		t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages\\user\\login.html"))
		t.Execute(w, "")
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	user, _ := dao.CheckUsername(username)
	if user.ID > 0 {
		t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages\\user\\regist.html"))
		t.Execute(w, "user already exists")
	}else {
		dao.SaveUser(username, password, email)
		t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages\\user\\login.html"))
		t.Execute(w, "")
	}
}