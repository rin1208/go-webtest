package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("kemo.html")
	t.Execute(w, "")
}
func kemomimi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "狐耳")
}

func main() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", hello)
	http.HandleFunc("/kemomimi", kemomimi)

	server.ListenAndServe()
}
