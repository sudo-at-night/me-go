package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sudo-at-night/me-go/src/parse"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html, err := parse.GetIndexContent(r.URL.Path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, html)
}

func main() {
	go parse.CacheIndex()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
