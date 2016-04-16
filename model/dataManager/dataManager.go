package dataManager

import (
//	"github.com/krix38/ScorchedGo/model/entity"
	"log"
)

type Action int

const (
	CREATE Action = iota
	READ
	UPDATE
	DELETE
)

type EntityAction struct {
	Entity interface{}
	Action Action
}

var RoomAction    = make(chan EntityAction)
var PlayerAction  = make(chan EntityAction)

func dispatchPlayerAction(action EntityAction){
	log.Println("triggered player action") /* TODO */
}

func dispatchRoomAction(action EntityAction){
	log.Println("triggered room action")   /* TODO */
}

func RunDataManager() {
//	roomList   := new(entity.RoomsList)
//	playerList := new(entity.PlayersList)
	for {
		select {
			case action := <-RoomAction:
				dispatchRoomAction(action)
			case action := <-PlayerAction:
				dispatchPlayerAction(action)
		}		
	}
}
