package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello world</h1>")
}

func contact(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<address>To get support, please contact to my <a href=\"mailto:linnaiwei@gmail.com\">email</a></address>")
}

func faq(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>This is my FAQ page</h1>")
}

func notFoundPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Wo could not find the page you were going >_< </h1>")
}

func main() {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(notFoundPage)
	router.GET("/", home)
	router.GET("/contact", contact)
	router.GET("/faq", faq)
	http.ListenAndServe(":3000", router)
}
