package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Pong")
	})

	http.ListenAndServe(":50059", nil)
}
