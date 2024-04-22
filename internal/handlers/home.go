package handlers

import (
	"errors"
	"net/http"
	"text/template"
)

// MainPageGetHandler handles GET requests to the main page.
func HomeHandler(w http.ResponseWriter, r *http.Request) error {

	logger.Info("User opens home page")

	// parse home page template
	tmpl := template.Must(template.ParseFiles(hopePageTemplateAdrress))

	errExceuteHompage := tmpl.Execute(w, nil)
	if errExceuteHompage != nil {
		errMsg = "----<-writeHomePageContent------<--ParseFiles----" + errExceuteHompage.Error()
		logMsg = "Error parsing  home page template"
		logger.Error(logMsg + errMsg)
		return errors.New(errMsg)
	} else {
		logger.Info(" Home Page template executed successfully")
	}
	return nil
}
