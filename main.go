package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello world</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<address>To get support, please contact to my <a href=\"mailto:linnaiwei@gmail.com\">email</a></address>")

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}
