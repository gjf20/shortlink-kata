package main

import (
	"log"
	"net/http"

	"example.com/shortlink-kata/shortlink" //todo change the repetitive name
)

func main() {
	mux := shortlink.GetMux()

	err := http.ListenAndServe(":8090", mux)
	log.Fatal(err)
}
