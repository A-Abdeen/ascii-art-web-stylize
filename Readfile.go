package asciiart

import (
	"html/template"
	"log"
	"net/http"
)

func Readfile(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/template.html")
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, "template.html")
}
