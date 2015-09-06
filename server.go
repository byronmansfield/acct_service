package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()

	log.Printf("Starting Server and listening at localhost%s", PORT)
	http.ListenAndServe(PORT, router)
}
