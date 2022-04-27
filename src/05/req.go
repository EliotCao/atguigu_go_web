package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "path ", r.URL.Path)
	fmt.Fprintln(w, "RawPath ", r.URL.RawPath)
	fmt.Fprintln(w, "RawQuery ", r.URL.RawQuery)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}