package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", Index)
	http.ListenAndServe(":8080", r)
}
