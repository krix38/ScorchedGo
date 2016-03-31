package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/krix38/ScorchedGo/properties"
	"github.com/krix38/ScorchedGo/web/controller/handler/page"
)

func registerStaticFileHandler(staticFilesUrl string) {
	http.Handle(staticFilesUrl, http.StripPrefix(
		staticFilesUrl, http.FileServer(
			http.Dir(
				properties.Configuration.StaticFilesPath))))
}

func registerMainPageHandler(mainHandlerUrl string) {
	http.HandleFunc(mainHandlerUrl, page.Main)
}

func RunWebController() {

	registerStaticFileHandler("/static/")
	registerMainPageHandler("/")

	//routing := make(map[string]func(http.ResponseWriter, *http.Request))

	log.Fatal(http.ListenAndServe(properties.Configuration.Port, context.ClearHandler(http.DefaultServeMux)))
}
