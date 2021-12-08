package foo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type person struct {
	Name string
	Age  int
}

var tom *person = &person{
	Name: "Tom",
	Age:  28,
}

func Controller(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		j, _ := json.Marshal(tom)
		w.Write(j)
	case "POST":
		d := json.NewDecoder(r.Body)
		p := &person{}
		err := d.Decode(p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		tom = p
		fmt.Fprint(w, "ok")
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "I can't do that.")
	}
}
