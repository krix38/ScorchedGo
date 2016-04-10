package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/krix38/ScorchedGo/properties"
	"github.com/krix38/ScorchedGo/web/controller/handler/page"
	"github.com/krix38/ScorchedGo/web/controller/handler/rest"
	"github.com/krix38/ScorchedGo/web"
)

func RunWebController() {

	web.RegisterStaticFileHandler("/static/")

	routing := make(map[string]web.Handler)

	routing["/"]                         =    web.Handler{HandlerFunc: page.Main, Session: false}
	routing["/api/getConnectionStatus"]  =    web.Handler{HandlerFunc: rest.GetConnectionStatus, Session: false}
	routing["/api/getAllChannels"]       =    web.Handler{HandlerFunc: rest.GetAllChannels, Session: true}
	
	web.MapHandlers(routing)

	log.Fatal(http.ListenAndServe(properties.Configuration.Port, context.ClearHandler(http.DefaultServeMux)))
}
