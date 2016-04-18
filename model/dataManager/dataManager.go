package dataManager

import (
	"log"

	"github.com/krix38/ScorchedGo/model/entity"
)

type Action int

const (
	CREATE Action = iota
	READ
	UPDATE
	DELETE
)

type EntityAction struct {
	Action         Action
	ResponseChan   chan interface{}
	Entity         interface{}
	AdditionalData interface{}
}

type sharedData struct {
	roomsId   int64
	rooms     entity.RoomsList
	playersId int64
	players   entity.PlayersList
}

var RoomAction   = make(chan EntityAction)
var PlayerAction = make(chan EntityAction)

func dispatchPlayerAction(action EntityAction, memData *sharedData) {
	log.Println("triggered player action") /* TODO */
}

func roomRead(action EntityAction, memData *sharedData) {
	if action.AdditionalData == nil { /* no filters */
		action.ResponseChan <- memData.rooms
	} /* TODO: filter response */
}

func roomCreate(action EntityAction, memData *sharedData) {
	room, ok := action.Entity.(entity.Room)
	if ok {
		memData.roomsId += 1
		room.Id = memData.roomsId
		memData.rooms.Rooms = append(memData.rooms.Rooms, room)
	} /* TODO: if not ok */
}

func dispatchRoomAction(action EntityAction, memData *sharedData) {
	switch action.Action {

	case CREATE:
		roomCreate(action, memData)
	case READ:
		roomRead(action, memData)
	case UPDATE:
	case DELETE:

	}
}

func RunDataManager() {
	memData := sharedData{
		rooms:     entity.RoomsList{Rooms: []entity.Room{}},
		roomsId:   0,
		players:   entity.PlayersList{Players: []entity.Player{}},
		playersId: 0,
	}
	for {
		select {

		case action := <-RoomAction:
			dispatchRoomAction(action, &memData)
		case action := <-PlayerAction:
			dispatchPlayerAction(action, &memData)

		}
	}
}
