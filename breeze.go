package main

import (
	"log"
	"net/http"
)

func main() {
	log.Print("server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
