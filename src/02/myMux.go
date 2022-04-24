package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "by myMux: ", r.URL.Path)
	})

	http.ListenAndServe(":8081", mux)
}
