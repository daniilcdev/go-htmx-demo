package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
)

type Film struct {
	Title    string
	Director string
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static")).ServeHTTP(w, r)
}

func postClicked(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./static/clicked.html"))

	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}

	tmpl.Execute(w, films)
}

func main() {
	handleSigTerms()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/clicked", postClicked)

	fmt.Println("Start listening...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleSigTerms() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("received SIGTERM, exiting")
		os.Exit(1)
	}()
}
