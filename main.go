package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

//Index  This function will return a welcome page
func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "<h1>Welcome! My friend ^o^</h1>")
}

//Hello  This function will say hello to the name you write in the named parameters
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "<h1>Hello! ", ps.ByName("name") ," :) </h1>")
}

func main() {
	router := httprouter.New()
    router.GET("/", Index)
    router.GET("/hello/:name", Hello)
	http.ListenAndServe(":3000", router)
}
