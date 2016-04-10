package rest

import (
	"github.com/krix38/ScorchedGo/model/api"
	"github.com/krix38/ScorchedGo/web/restFactory"
)

var ConnectionStatus = restFactory.CreateRestHandler(connectionStatus, []string{"GET"}, nil)
var GetAllChannels   = restFactory.CreateRestHandler(getAllChannels, []string{"GET"}, nil)

func connectionStatus(restData *restFactory.RestHandlerData) (interface{}, error) {
	/* TODO: check connection status */
	return api.ConnectionInfo{SignedIn: false}, nil
}

func getAllChannels(restData *restFactory.RestHandlerData) (interface{}, error) {
	/* TODO: get all channels */
	return api.ChannelsList{Channels: []api.Channel{}}, nil
}