package controller

import (
	"fmt"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}
