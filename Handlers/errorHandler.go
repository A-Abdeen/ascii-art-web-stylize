package asciiart

import (
	mod "asciiart"
	"html/template"
	"log"
	"net/http"
)

func ErrorHandler(w http.ResponseWriter, r *http.Request, statusCode int) {
	var errorMsg string
	files := []string{
		"templates/mainpage.html",
		"templates/error.tmpl",
	}

	switch {
	case statusCode == 404:
		errorMsg = "Page Not Found"
	case statusCode == 405:
		errorMsg = "Method Not Allowed"
	case statusCode == 400:
		errorMsg = "Bad User Request"
	case statusCode == 500:
		errorMsg = "Internal Server Error"
	case statusCode == 418:
		errorMsg = "The Server Refuses The Attempt To Brew Coffee With A Teapot\r\n(Use english letters only!)"
	default:
		errorMsg = "UNFAMILIAR ERROR WHAT DIS :("
	}

	response := mod.Err{
		Status:    true,
		ErrorMsg:  errorMsg,
		ErrorCode: statusCode,
		Year:      2023,
	}

	w.WriteHeader(response.ErrorCode)
	t, err := template.ParseFiles(files...)
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(w, "mainpage.html", response)
}
