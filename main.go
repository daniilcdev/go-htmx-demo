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
	tmpl := template.Must(template.ParseFiles("./static/index.html"))

	films := map[string][]Film{
		"Films": {
			{Title: "The Godfather", Director: "Francis Ford Coppola"},
			{Title: "Blade Runner", Director: "Ridley Scott"},
			{Title: "The Thing", Director: "John Carpenter"},
		},
	}

	tmpl.Execute(w, films)
}

func addFilmHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	director := r.PostFormValue("director")
	htmlString := fmt.Sprintf("<li class='list-group-item bg-primary text-white'>%s - %s</li>", title, director)
	tmpl, _ := template.New("t").Parse(htmlString)

	tmpl.Execute(w, nil)
}

func main() {
	handleSigTerms()

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add-film/", addFilmHandler)

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
