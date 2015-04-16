package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var port int

	flag.IntVar(&port, "port", 8080, "port to listen to...")
	flag.Parse()

	input := strconv.Itoa(port)

	log.Print("server listening on port " + input + ".")
	if err := http.ListenAndServe(":"+input, nil); err != nil {
		log.Fatal(err)
	}
}
