package asciiart

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

type Art struct {
	Output string
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
	if IsAscii(inputString) {
		banner := r.FormValue("banner")
		output := Art{Output: AsciiArt(inputString, banner)}
		if strings.HasPrefix(output.Output, "500: ") {
			http.Error(w, "Internal Server Error: "+strings.TrimPrefix(output.Output, "500: "), http.StatusInternalServerError)
			return
		}
		if strings.HasPrefix(output.Output, "400: ") {
			http.Error(w, "Bad Request: "+strings.TrimPrefix(output.Output, "400: "), http.StatusBadRequest)
			return
		}
		t, err := template.ParseFiles("/Users/abdyn/Codes/ascii-art-web-stylize/webprogram/templates/template.html")
		if err != nil {
			log.Fatal(err)
		}
		t.Execute(w, output)
	} else {
		http.Error(w, "Invalid input: Only use English letters, numbers & characters", http.StatusBadRequest)
	}
}
