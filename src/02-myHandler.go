package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {}

func (m *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "by MyHandler: " + request.URL.Path)
}

func main() {
	myHandler := MyHandler{}
	http.Handle("/myHandler", &myHandler)
	http.ListenAndServe(":8081",nil)
}
