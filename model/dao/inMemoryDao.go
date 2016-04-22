package dao

import (
	"github.com/krix38/ScorchedGo/model/dataManager"
	"github.com/krix38/ScorchedGo/model/entity"
)

func CreateRoom(room entity.Room) error {
	responseChan := make(chan interface{})
	dataManager.RoomAction <- dataManager.EntityAction{
		Entity: room,
		Action: dataManager.CREATE,
	}
	return (<-responseChan).(error)
}

func CreatePlayer(player entity.Player) error {
	responseChan := make(chan interface{})
	dataManager.PlayerAction <- dataManager.EntityAction{
		Entity: player,
		ResponseChan: responseChan,
		Action: dataManager.CREATE,
	}
	return (<-responseChan).(error)
}

func LoadRooms() *entity.RoomsList {
	responseChan := make(chan interface{})
	dataManager.RoomAction <- dataManager.EntityAction{
		ResponseChan: responseChan,
		Action: dataManager.READ,
	}
	rooms := <-responseChan
	roomsConverted, ok := rooms.(entity.RoomsList)
	if ok {
		return &roomsConverted
	}else{
		return nil
	}
}
