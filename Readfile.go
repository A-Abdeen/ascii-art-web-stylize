package asciiart

import (
	"html/template"
	"log"
	"net/http"
)

// type Example struct {
// 	Standard string
// 	Shadow string
// 	Thinkertoy string
// 	Input string	
// }

func Readfile(w http.ResponseWriter, r *http.Request) {
Examples := Art{Standard: AsciiArt("standard", "standard"), Shadow: AsciiArt("shadow", "shadow"),
	 Thinkertoy: AsciiArt("thinkertoy", "thinkertoy"), Input: "", Color: "#9141ac", BackColor: "#a7bfc2"}
	t, err := template.ParseFiles("templates/template.html")
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(w, "template.html", Examples)
}
