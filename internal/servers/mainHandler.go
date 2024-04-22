package servers

import (
	//"errors"
	//"log"

	"net/http"

	"github.com/ediallocyf/asciiartweb/internal/handlers"
)

// MainHandler is the main HTTP request handler.
func MainHandler(w http.ResponseWriter, r *http.Request) {

	switch {
	case r.URL.Path == homePagePath && r.Method == http.MethodGet:

		errHomeHandler := handlers.HomeHandler(w, r)
		if errHomeHandler != nil {
			errMsg := "--MainHandler--<------<---HomeHandler---: " + errHomeHandler.Error()
			logMsg := "Error in loading home page"
			logger.Error(logMsg + errMsg)
			return
		}
		logger.Info("Home page loaded successfully")

		//--------------------------------------------------------------
	case r.URL.Path == ascciArtPagePath && r.Method == http.MethodPost:
		errAsciiArt := handlers.GenerateAsciiArtHandler(w, r)
		if errAsciiArt != nil {
			errMsg := "--MainHandler--<------<---GenerateAsciiArtHandler---: " + errAsciiArt.Error()
			logMsg := "Error in loading ASCII art page"
			logger.Error(logMsg + errMsg)
			return
		}
		logger.Info("ASCII art page loaded successfully")

		//-------------------------------------------------------------------
	default:
		logger.Info("User tried to access an unknown page")
		handlers.ErrorPage(w, NotFound404)
	}
}
