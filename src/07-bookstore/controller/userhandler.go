package controller

import (
	"atguigu_go_web/src/07-bookstore/dao"
	"atguigu_go_web/src/07-bookstore/model"
	util "atguigu_go_web/src/07-bookstore/utils"
	"html/template"
	"net/http"
)

//Login handle user login
func Login(w http.ResponseWriter, r *http.Request) {
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	user, _ := dao.CheckUsernameAndPassword(username, password)
	if user.ID > 0 {
		uuid := util.CreateUUID()
		session := &model.Session{
			uuid,
			user.Username,
			user.ID,
		}
		dao.AddSession(session)
		cookie := http.Cookie{
			Name: "uuid",
			Value: uuid,
			HttpOnly: true,
		}
		//r.Header.Set("Set-Cookie", cookie.String())
		http.SetCookie(w, &cookie)

		t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages\\user\\login_success.html"))
		t.Execute(w, "")
	}else {
		t := template.Must(template.ParseFiles("C:\\Users\\RZNQGT\\Desktop\\atguigu_go_web\\src\\07-bookstore\\views\\pages\\user\\login.html"))
		t.Execute(w, "")
	}
}

// Logout handle user logout
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		cookieValue := cookie.Value
		dao.DeleteSession(cookieValue)
		cookie.MaxAge = -1
		http.SetCookie(w, cookie)
	}
	//goto index
	GetPageBooksByPrice(w, r)
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

//CheckUsername validate username valid by send ajax request
func CheckUsername(w http.ResponseWriter, r *http.Request)  {
	username:= r.PostFormValue("username")
	user, _ := dao.CheckUsername(username)
	if user.ID > 0 {
		w.Write([]byte("user already exists"))
	}else {
		w.Write([]byte("user valid"))
	}
}