package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MagnusFrater/helmet"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "I love HelmetJS, I wish there was a Go(lang) equivalent...")
	})

	helmet := helmet.Default()
	http.Handle("/", helmet.Secure(handler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
