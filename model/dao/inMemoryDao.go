package dao

import (
	"github.com/krix38/ScorchedGo/model/dataManager"
	"github.com/krix38/ScorchedGo/model/entity"
)

func CreateRoom(room entity.Room) error {
	responseChan := make(chan interface{})
	dataManager.RoomAction <- dataManager.EntityAction{
		Entity: room,
		ResponseChan: responseChan,
		Action: dataManager.CREATE,
	}
	err := <-responseChan
	if(err != nil){
		return err.(error)
	}else{
		return nil
	}
	
}

func CreatePlayer(player entity.Player) error {
	responseChan := make(chan interface{})
	dataManager.PlayerAction <- dataManager.EntityAction{
		Entity:       player,
		ResponseChan: responseChan,
		Action:       dataManager.CREATE,
	}
	err := <-responseChan
	if(err != nil){
		return err.(error)
	}else{
		return nil
	}
}

func LoadRooms() *entity.RoomsList {
	responseChan := make(chan interface{})
	dataManager.RoomAction <- dataManager.EntityAction{
		ResponseChan: responseChan,
		Action:       dataManager.READ,
		AdditionalData: dataManager.ReadInfo{
			Mode: dataManager.ALL,
		},
	}
	rooms := <-responseChan
	roomsConverted, ok := rooms.(entity.RoomsList)
	if ok {
		return &roomsConverted
	} else {
		return nil
	}
}

func UpdateRoom(room entity.Room) error {
	responseChan := make(chan interface{})
	dataManager.RoomAction <- dataManager.EntityAction{
		Entity:       room,
		ResponseChan: responseChan,
		Action:       dataManager.UPDATE,
	}
	return (<-responseChan).(error)
}
