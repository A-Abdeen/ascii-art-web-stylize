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
		mainerror = err
		er, err := template.ParseFiles("templates/error")
		if err != nil {
			log.Fatal(err)
		}
		er.Execute(w, mainerror)
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
			// http.Error(w, "Internal Server Error: "+strings.TrimPrefix(output.Output, "500: "), http.StatusInternalServerError)
			// return

		}
		if strings.HasPrefix(output.Output, "400: ") {
			http.Error(w, "Bad Request: "+strings.TrimPrefix(output.Output, "400: "), http.StatusBadRequest)
			return
		}
		t, err := template.ParseFiles("templates/template.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, output)
	} else {
		http.Error(w, "Invalid input: Only use English letters, numbers & characters", http.StatusBadRequest)
	}
}

