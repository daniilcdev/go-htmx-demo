package main

import (
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static")).ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":3000", nil)
}
