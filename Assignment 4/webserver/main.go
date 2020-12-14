package main

import (
	"log"
	"net/http"
)

func main() {

	http.Handle("/", http.FileServer(http.Dir("./divya")))

	log.Fatal(http.ListenAndServe(":9090", nil))
}