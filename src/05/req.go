package main

import (
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

	err := r.ParseForm()
	fmt.Println(err)
	fmt.Fprintln(w, "请求参数", r.Form)
	fmt.Fprintln(w, "请求参数", r.PostForm)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}