package rest

import (
	"github.com/krix38/ScorchedGo/model/dao"
	"github.com/krix38/ScorchedGo/model/entity"
	"github.com/krix38/ScorchedGo/web/restFactory"
	"errors"
)

var GetConnectionStatus = restFactory.CreateRestHandler(getConnectionStatus, []string{"GET"},    nil)
var GetAllRooms         = restFactory.CreateRestHandler(getAllRooms,         []string{"GET"},    nil)
var CreateRoom          = restFactory.CreateRestHandler(createRoom,          []string{"CREATE"}, entity.Room{})

func getConnectionStatus(restData *restFactory.RestHandlerData) (interface{}, error) {
	/* TODO: check connection status */
	return entity.ConnectionInfo{SignedIn: false}, nil
}

func getAllRooms(restData *restFactory.RestHandlerData) (interface{}, error) {
	rooms := dao.LoadRooms()
	if rooms != nil {
		return rooms, nil
	}else{
		return nil, errors.New("error loading rooms list")
	}
}

func createRoom(restData *restFactory.RestHandlerData) (interface{}, error) {
	room, ok := restData.InputJsonObj.(entity.Room)
	if ok {
		err := dao.CreateRoom(room)
		if err != nil {
			return entity.Msg{ Message: "room created" }, nil
		}else{
			return nil, err
		}
	}
	return nil, errors.New("error creating room")
}
