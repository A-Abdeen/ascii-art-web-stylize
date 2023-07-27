package main

import (
	handler "asciiart/Handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler.MainHandler)
	http.HandleFunc("/asciiart", handler.ArtHandler)
	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
