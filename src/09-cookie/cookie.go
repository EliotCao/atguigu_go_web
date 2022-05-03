package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name: "user",
		Value: "admin",
		HttpOnly: true,
		MaxAge: 60,
	}
	//w.Header().Set("Set-cookie", cookie.String())
	http.SetCookie(w, &cookie)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	cookie := r.Header["Cookie"]
	fmt.Println(cookie)
	cookie2, _ := r.Cookie("user")
	fmt.Println(cookie2)
}

func main() {
	http.HandleFunc("/cookie", setCookie)
	http.HandleFunc("/getcookie", getCookie)
	http.ListenAndServe(":8081", nil)
}