package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static")).ServeHTTP(w, r)
}

func postClicked(w http.ResponseWriter, r *http.Request) {
	div, e := os.ReadFile("./static/clicked.html")
	if e != nil {
		panic(e)
	}

	io.WriteString(w, string(div[:]))
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
