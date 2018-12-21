package main 

import (
	"fmt"
	"net/http"
)

func hello (w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "ケモミミ")
	}
func kemomimi(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "狐耳")
    }
    
func main () {	

	server := http.Server {
	Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello",  hello)
	http.HandleFunc("/kemomimi",  kemomimi)
	

	server.ListenAndServe()
}

