package dataManager

import (
	"errors"

	"github.com/krix38/ScorchedGo/model/entity"
)

type Action int

const (
	CREATE Action = iota
	READ
	UPDATE
	DELETE
)

type ReadMode int

const (
	SINGLE ReadMode = iota
	ALL
)

type ReadInfo struct {
	Mode  ReadMode
	Index int64
}

type EntityAction struct {
	Action         Action
	ResponseChan   chan interface{}
	Entity         interface{}
	AdditionalData ReadInfo
}

type sharedData struct {
	roomsId   int64
	rooms     entity.RoomsList
	playersId int64
	players   entity.PlayersList
}

var RoomAction = make(chan EntityAction)
var PlayerAction = make(chan EntityAction)

func findRoom(id int64, memData *sharedData) (int, *entity.Room) {
	for index, room := range memData.rooms.Rooms {
		if room.Id == id {
			return index, &room
		}
	}
	return -1, nil
}

func findPlayer(id int64, memData *sharedData) (int, *entity.Player) {
	for index, player := range memData.players.Players {
		if player.Id == id {
			return index, &player
		}
	}
	return -1, nil
}

func roomCreate(action EntityAction, memData *sharedData) {
	room, ok := action.Entity.(entity.Room)
	if ok {
		memData.roomsId += 1
		room.Id = memData.roomsId
		memData.rooms.Rooms = append(memData.rooms.Rooms, room)
		action.ResponseChan <- nil
	} else {
		action.ResponseChan <- errors.New("failed to create room")
	}
}

func roomRead(action EntityAction, memData *sharedData) {
	switch action.AdditionalData.Mode {
		
	case ALL:
		action.ResponseChan <- memData.rooms
	case SINGLE:
		_, room := findRoom(action.AdditionalData.Index, memData)
		action.ResponseChan <- room
	}
}

func roomUpdate(action EntityAction, memData *sharedData) {
	room, ok := action.Entity.(entity.Room)
	if ok {
		index, _ := findRoom(room.Id, memData)
		memData.rooms.Rooms[index] = room
		action.ResponseChan <- nil
	} else {
		action.ResponseChan <- errors.New("failed to update room")
	}
}

func roomDelete(action EntityAction, memData *sharedData) {
	room, ok := action.Entity.(entity.Room)
	if ok {
		index, _ := findRoom(room.Id, memData)
		memData.rooms.Rooms = append(memData.rooms.Rooms[:index], memData.rooms.Rooms[index+1:]...)
		action.ResponseChan <- nil
	} else {
		action.ResponseChan <- errors.New("failed to delete room")
	}
}

func playerCreate(action EntityAction, memData *sharedData) {
	player, ok := action.Entity.(entity.Player)
	if ok {
		memData.playersId += 1
		player.Id = memData.playersId
		memData.players.Players = append(memData.players.Players, player)
		action.ResponseChan <- nil
	} else {
		action.ResponseChan <- errors.New("failed to create player")
	}
}

func playerRead(action EntityAction, memData *sharedData) {
	switch action.AdditionalData.Mode {
		
	case ALL:
		action.ResponseChan <- memData.players
	case SINGLE:
		_, player := findPlayer(action.AdditionalData.Index, memData)
		action.ResponseChan <- player
	}
}

func playerUpdate(action EntityAction, memData *sharedData) {
	player, ok := action.Entity.(entity.Player)
	if ok {
		index, _ := findPlayer(player.Id, memData)
		memData.players.Players[index] = player
		action.ResponseChan <- nil
	} else {
		action.ResponseChan <- errors.New("failed to update player")
	}
}

func playerDelete(action EntityAction, memData *sharedData) {
	player, ok := action.Entity.(entity.Player)
	if ok {
		index, _ := findPlayer(player.Id, memData)
		memData.players.Players = append(memData.players.Players[:index], memData.players.Players[index+1:]...)
		action.ResponseChan <- nil
	} else {
		action.ResponseChan <- errors.New("failed to delete player")
	}
}

func dispatchRoomAction(action EntityAction, memData *sharedData) {
	switch action.Action {

	case CREATE:
		roomCreate(action, memData)
	case READ:
		roomRead(action, memData)
	case UPDATE:
		roomUpdate(action, memData)
	case DELETE:
		roomDelete(action, memData)
	}
}

func dispatchPlayerAction(action EntityAction, memData *sharedData) {
	switch action.Action {

	case CREATE:
		playerCreate(action, memData)
	case READ:
		playerRead(action, memData)
	case UPDATE:
		playerUpdate(action, memData)
	case DELETE:
		playerDelete(action, memData)
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
