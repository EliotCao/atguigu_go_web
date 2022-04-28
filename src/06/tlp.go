package main

import "net/http"

func template(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location","https://www.baidu.com")
	w.WriteHeader(302)
}

func main() {
	http.HandleFunc("/template", template)
	http.ListenAndServe(":8081", nil)
}
