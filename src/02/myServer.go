package main

import (
	"fmt"
	"net/http"
	"time"
)

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "by MyHandler: "+request.URL.Path)
}

func main() {
	myhandler := MyHandler{}
	server := http.Server{
		Addr:        ":8081",
		Handler:     &myhandler,
		ReadTimeout: 1 * time.Second,
	}
	server.ListenAndServe()
}
