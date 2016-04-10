package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/krix38/ScorchedGo/properties"
	"github.com/krix38/ScorchedGo/web/controller/handler/page"
	"github.com/krix38/ScorchedGo/web/controller/handler/rest"
	"github.com/krix38/ScorchedGo/web/session"
)

type handler struct {
	handlerFunc func(http.ResponseWriter, *http.Request)
	session     bool
}

func RunWebController() {

	registerStaticFileHandler("/static/")

	routing := make(map[string]handler)

	routing["/"]                         =    handler{handlerFunc: page.Main, session: false}
	routing["/api/getConnectionStatus"]  =    handler{handlerFunc: rest.GetConnectionStatus, session: false}
	routing["/api/getAllChannels"]       =    handler{handlerFunc: rest.GetAllChannels, session: true}
	
	mapHandlers(routing)

	log.Fatal(http.ListenAndServe(properties.Configuration.Port, context.ClearHandler(http.DefaultServeMux)))
}

func registerStaticFileHandler(staticFilesUrl string) {
	http.Handle(staticFilesUrl, http.StripPrefix(
		staticFilesUrl, http.FileServer(
			http.Dir(
				properties.Configuration.StaticFilesPath))))
}

func mapHandlers(routingMap map[string]handler) {
	for url, handler := range routingMap {
		if handler.session {
			session.CreateSessionHandler(url, handler.handlerFunc)
		} else {
			http.HandleFunc(url, handler.handlerFunc)
		}
	}
}