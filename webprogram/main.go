package main

import (
	"asciiart"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", asciiart.Readfile)
	http.HandleFunc("/asciiart", asciiart.ArtHandler)
	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
