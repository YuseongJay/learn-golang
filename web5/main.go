package main

import (
	"fmt"
	"net/http"
)

type fooHandler struct{}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello Food!</h1>`")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>My First Heading</h1> <p>My first paragraph.</p>")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/bar", barHandler)
	mux.Handle("/foo", &fooHandler{})
	http.ListenAndServe(":80", mux)
}
