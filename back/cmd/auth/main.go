package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/register", handler)
	http.HandleFunc("/login", handler)
	http.HandleFunc("/refresh-token", handler)
	log.Print("Starting...")
	
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}