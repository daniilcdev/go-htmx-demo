package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static")).ServeHTTP(w, r)
}

func main() {
	handleSigTerms()

	http.HandleFunc("/", indexHandler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println("http.ListenAndServe:", err)
		os.Exit(1)
	}
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
