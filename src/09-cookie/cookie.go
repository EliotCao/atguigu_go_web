package main

import "net/http"

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name: "user",
		Value: "admin",
		HttpOnly: true,
	}
	//w.Header().Set("Set-cookie", cookie.String())
	http.SetCookie(w, &cookie)
}

func main() {
	http.HandleFunc("/cookie", setCookie)
	http.ListenAndServe(":8081", nil)
}