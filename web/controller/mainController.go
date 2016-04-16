package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/context"
	"github.com/krix38/ScorchedGo/properties"
	"github.com/krix38/ScorchedGo/web"
	"github.com/krix38/ScorchedGo/web/controller/handler/page"
	"github.com/krix38/ScorchedGo/web/controller/handler/rest"
)

func RunWebController() {

	web.RegisterStaticFileHandler("/static/")

	routing := make(map[string]web.Handler)

	routing["/"]                          =  web.Handler{ HandlerFunc: page.Main,                 Session: false }
	routing["/api/connectionStatus/get"]  =  web.Handler{ HandlerFunc: rest.GetConnectionStatus,  Session: false }
	routing["/api/room/create"]           =  web.Handler{ HandlerFunc: rest.CreateRoom,           Session: false }
	routing["/api/allChannels/get"]       =  web.Handler{ HandlerFunc: rest.GetAllRooms,          Session: true  }
	
	web.MapHandlers(routing)

	log.Fatal(http.ListenAndServe(properties.Configuration.Port, context.ClearHandler(http.DefaultServeMux)))
}
