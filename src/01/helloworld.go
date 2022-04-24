package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, request.URL.Path)
	})
	http.ListenAndServe(":8081",nil)
}
