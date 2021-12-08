package bar

import (
	"fmt"
	"net/http"
)

func Controller(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "localhost/bar?name=\"message!\""
	}
	fmt.Fprintf(w, "%s", name)
}
