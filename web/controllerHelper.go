package web

import (
	"net/http"

	"github.com/krix38/ScorchedGo/properties"
	"github.com/krix38/ScorchedGo/web/session"
)

type Handler struct {
	HandlerFunc func(http.ResponseWriter, *http.Request)
	Session     bool
}

func RegisterStaticFileHandler(staticFilesUrl string) {
	http.Handle(staticFilesUrl, http.StripPrefix(
		staticFilesUrl, http.FileServer(
			http.Dir(
				properties.Configuration.StaticFilesPath))))
}

func MapHandlers(routingMap map[string]Handler) {
	for url, handler := range routingMap {
		if handler.Session {
			session.CreateSessionHandler(url, handler.HandlerFunc)
		} else {
			http.HandleFunc(url, handler.HandlerFunc)
		}
	}
}
