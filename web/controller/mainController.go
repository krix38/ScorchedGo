package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/krix38/ScorchedGo/properties"
	"github.com/krix38/ScorchedGo/web/controller/handler/page"
)

func registerStaticFileHandler() {
	staticFilesUrl := "/static/"
	http.Handle(staticFilesUrl, http.StripPrefix(
		staticFilesUrl, http.FileServer(
			http.Dir(
				properties.Configuration.StaticFilesPath))))
}

func registerMainPageHandler() {
	http.HandleFunc("/", page.Main)
}

func RunWebController() {

	registerStaticFileHandler()

	registerMainPageHandler()

	//routing := make(map[string]func(http.ResponseWriter, *http.Request))

	log.Fatal(http.ListenAndServe(properties.Configuration.Port, context.ClearHandler(http.DefaultServeMux)))
}
