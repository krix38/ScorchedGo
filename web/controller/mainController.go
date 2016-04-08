package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/krix38/ScorchedGo/properties"
	"github.com/krix38/ScorchedGo/web/controller/handler/page"
	"github.com/krix38/ScorchedGo/web/controller/handler/rest"
)

type handler struct {
	handlerFunc func(http.ResponseWriter, *http.Request)
	session     bool
}

func RunWebController() {

	registerStaticFileHandler("/static/")

	routing := make(map[string]handler)

	routing["/"]                         =    handler{handlerFunc: page.Main, session: false}
	routing["/api/connectionStatus"]     =    handler{handlerFunc: rest.ConnectionStatus, session: false}
	//
	routing["/api/test"] = handler{handlerFunc: rest.TestFunc, session: false}
	//

	mapHandlers(routing)

	log.Fatal(http.ListenAndServe(properties.Configuration.Port, context.ClearHandler(http.DefaultServeMux)))
}

func registerStaticFileHandler(staticFilesUrl string) {
	http.Handle(staticFilesUrl, http.StripPrefix(
		staticFilesUrl, http.FileServer(
			http.Dir(
				properties.Configuration.StaticFilesPath))))
}

func createSessionHandler(url string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(url, handler)
}

func mapHandlers(routingMap map[string]handler) {
	for url, handler := range routingMap {
		if handler.session {
			createSessionHandler(url, handler.handlerFunc)
		} else {
			http.HandleFunc(url, handler.handlerFunc)
		}
	}
}
