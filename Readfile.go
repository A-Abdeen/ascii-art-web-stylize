package asciiart

import (
	"html/template"
	"log"
	"net/http"
)

type Example struct {
	Standard string
	Shadow string
	Thinkertoy string
	Input string	
}

func Readfile(w http.ResponseWriter, r *http.Request) {
	Examples := Example{Standard: AsciiArt("standard", "standard"), Shadow: AsciiArt("shadow", "shadow"),
	 Thinkertoy: AsciiArt("thinkertoy", "thinkertoy"), Input: ""}
	t, err := template.ParseFiles("templates/template.html")
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(w, "template.html", Examples)
}
