package main

import (
	"atguigu_go_web/src/05/model"
	json2 "encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "path ", r.URL.Path)
	fmt.Fprintln(w, "RawPath ", r.URL.RawPath)
	fmt.Fprintln(w, "RawQuery ", r.URL.RawQuery)
	fmt.Fprintln(w, "Header", r.Header)
	fmt.Fprintln(w, "Accept-encoding by [] ", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "Accept-encoding by get() ", r.Header.Get("Accept-Encoding"))
	fmt.Fprintln(w, "ContentLength ", r.ContentLength)
	//len := r.ContentLength
	//data := make([]byte, len)
	//r.Body.Read(data)
	//fmt.Fprintln(w, "body ", data)

	//err := r.ParseForm()
	//fmt.Println(err)
	//fmt.Fprintln(w, "请求参数", r.Form)
	//fmt.Fprintln(w, "请求参数", r.PostForm)
	fmt.Fprintln(w, "请求参数", r.FormValue("user"))
	fmt.Fprintln(w, "请求参数", r.PostFormValue("username"))
}

func jsonRsp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	user := model.User{
		ID: 1,
		Username:"json",
		Password:"123",
		Email:"test@go.com",
	}
	json,_ := json2.Marshal(user)
	w.Write(json)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/json", jsonRsp)
	http.ListenAndServe(":8081", nil)
}