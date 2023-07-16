package main

import (
	"asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Art struct {
	Output string
	mainerror error
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/asciiart", ArtHandler)
	fmt.Printf("starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func ArtHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if r.Method != "POST" || r.URL.Path != "/asciiart" {
		http.Error(w, "Only POST requests are allowed.\nUse the main page to create new art", http.StatusMethodNotAllowed)
		return
	}

	inputString := r.FormValue("input")
	if asciiart.IsAscii(inputString) {
		banner := r.FormValue("banner")
		output := Art{Output: asciiart.AsciiArt(inputString, banner)}
		if strings.HasPrefix(output.Output, "500: ") {
			http.ServeFile(w, r, "templates/error500.html")
		}
		if strings.HasPrefix(output.Output, "400: ") {
			http.ServeFile(w, r, "templates/error400.html")
		}
		t, err := template.ParseFiles("templates/template.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, output)
	} else {
			// fmt.Fprintf(w, "Invalid input: Only use English letters, numbers & characters")
			http.ServeFile(w, r, "templates/error400.html")	
	}
}

