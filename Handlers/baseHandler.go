package asciiart

import (
	mod "asciiart"
	"html/template"
	"log"
	"net/http"
	"os"
)

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/base.html",
		"templates/form.tmpl",
		"templates/output.tmpl",
		"templates/error.tmpl",
	}
	if r.Method != "GET" {
		ErrorHandler(w, r, http.StatusMethodNotAllowed)
		return
	} else if r.URL.Path != "/" {
		ErrorHandler(w, r, http.StatusNotFound)
		return
	} else {
		standardbanner,_ := os.ReadFile("standardbanner")
		shadowbanner,_ := os.ReadFile("shadowbanner")
		thinkertoybanner,_ := os.ReadFile("thinkertoybanner")
		examples := mod.Art{
			Standard:   string(standardbanner),
			Shadow:     string(shadowbanner),
			Thinkertoy: string(thinkertoybanner),
			Input:      "",
			Color:      "#9141ac",
			BackColor:  "#a7bfc2",
			Year:       2023,
		}
		t, err := template.ParseFiles(files...)
		if err != nil {
			log.Fatal(err)
		}
		t.ExecuteTemplate(w, "base.html", examples)
	}
}
