package rest

import (
	"github.com/krix38/ScorchedGo/model/dataManager"
	"github.com/krix38/ScorchedGo/model/entity"
	"github.com/krix38/ScorchedGo/web/restFactory"
)

var GetConnectionStatus = restFactory.CreateRestHandler(getConnectionStatus, []string{"GET"}, nil)
var GetAllRooms         = restFactory.CreateRestHandler(getAllRooms,         []string{"GET"}, nil)
var CreateRoom          = restFactory.CreateRestHandler(createRoom,          []string{"GET"}, nil)

func getConnectionStatus(restData *restFactory.RestHandlerData) (interface{}, error) {
	/* TODO: check connection status */
	return entity.ConnectionInfo{SignedIn: false}, nil
}

func getAllRooms(restData *restFactory.RestHandlerData) (interface{}, error) {
	/* TODO: get all channels */
	return entity.RoomsList{Rooms: []entity.Room{}}, nil
}

func createRoom(restData *restFactory.RestHandlerData) (interface{}, error) {
	dataManager.RoomAction <- dataManager.EntityAction{
		Entity: entity.Room{},
		Action: dataManager.CREATE,
	}
	return nil, nil
}
