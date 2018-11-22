package main 

import (
	"log"
	"net/http"
)

func main () {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte( `
	<!DOCTYPE html>
	<html lang = "ja">
	<head>
	<meta charset="UTF-8">
	<title> インフォ</title>
	<body>
	<h1>ようこそ</h1>
	</body>
	<html>
	`))
})

if err := http.ListenAndServe(":8080", nil); err !=nil {
	log.Fatal("ListenAndServe:", err)

}

}
