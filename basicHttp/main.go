package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server started...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to Go http service")
	})

	http.HandleFunc("/greet", Greet)
	http.HandleFunc("/health", Message("Ok"))
	http.HandleFunc("/ping", Message("pong"))
	http.ListenAndServe(":50080", nil)
}

func Message(message string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, message)
	}
}

func Greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
