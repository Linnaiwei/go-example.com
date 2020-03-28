package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Hello world</h1>")
	} else if r.URL.Path == "/contact" {
		fmt.Fprint(w, "<address>To get support, please contact to my <a href=\"mailto:linnaiwei@gmail.com\">email</a></address>")
	}
	
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
