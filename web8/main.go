package main

import (
	"net/http"
	"web8/controller"
	"web8/controller/bar"
	"web8/controller/foo"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.Controller)
	mux.HandleFunc("/bar", bar.Controller)
	mux.HandleFunc("/foo", foo.Controller)
	http.ListenAndServe(":80", mux)
}
