package asciiart

import (
	mod "asciiart"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func ArtHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/mainpage.html",
		"templates/form.tmpl",
		"templates/error.tmpl",
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	if r.Method != "POST" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	} else if r.URL.Path != "/asciiart" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	} else {
		inputString := r.FormValue("input")
		if mod.IsAscii(inputString) {
			banner := r.FormValue("banner")
			color := r.FormValue("color")
			backColor := r.FormValue("backColor")
			output := mod.Art{
				Output:     mod.AsciiArt(inputString, banner),
				Standard:   mod.AsciiArt("standard", "standard"),
				Shadow:     mod.AsciiArt("shadow", "shadow"),
				Thinkertoy: mod.AsciiArt("thinkertoy", "thinkertoy"),
				Input:      inputString,
				Color:      color,
				BackColor:  backColor,
				Year:       2023,
			}
			if strings.HasPrefix(output.Output, "500: ") {
				ErrorHandler(w, r, http.StatusInternalServerError)
				return
			}
			if strings.HasPrefix(output.Output, "400: ") {
				ErrorHandler(w, r, http.StatusBadRequest)
				return
			}
			t, err := template.ParseFiles(files...)
			if err != nil {
				log.Fatal(err)
			}
			t.ExecuteTemplate(w, "mainpage.html", output)
		} else {
			ErrorHandler(w, r, http.StatusTeapot)
			return
		}
	}
}
