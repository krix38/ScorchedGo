package controller

import (
	"log"
	"net/http"

	"github.com/krix38/ScorchedGo/external/github.com/gorilla/context"
	"github.com/krix38/ScorchedGo/properties"
)

func registerStaticFileHandler() {
	http.Handle(properties.Configuration.StaticFilesURL, http.StripPrefix(
		properties.Configuration.StaticFilesURL, http.FileServer(
			http.Dir(
				properties.Configuration.StaticFilesPath))))
}

func RunWebController() {
	registerStaticFileHandler()
	log.Fatal(http.ListenAndServe(properties.Configuration.Port, context.ClearHandler(http.DefaultServeMux)))
}
