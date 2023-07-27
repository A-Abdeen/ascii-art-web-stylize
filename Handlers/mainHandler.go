package asciiart

import (
	mod "asciiart"
	"html/template"
	"log"
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/mainpage.html",
		"templates/error.tmpl",
	}
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	} else if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	} else {
		examples := mod.Art{
			Standard:   mod.AsciiArt("standard", "standard"),
			Shadow:     mod.AsciiArt("shadow", "shadow"),
			Thinkertoy: mod.AsciiArt("thinkertoy", "thinkertoy"),
			Input:      "",
			Color:      "#9141ac",
			BackColor:  "#a7bfc2",
			Year:       2023,
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			log.Fatal(err)
		}
		t.ExecuteTemplate(w, "mainpage.html", examples)
	}
}
