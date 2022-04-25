package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "Test HTPP")
}

func main() {
	
}
