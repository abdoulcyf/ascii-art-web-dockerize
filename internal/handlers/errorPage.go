package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func ErrorPage(w http.ResponseWriter, errMsg string) {
	// Define the template
	t := template.Must(template.ParseFiles(errorTemplAdrr))
	data := struct {
		Msg string
		Url string
	}{
		Msg: errMsg,
		Url: homePagePath,
	}
	err := t.Execute(w, data)
	if err != nil {
		log.Printf("Failed to execute error page template")
	}
}
