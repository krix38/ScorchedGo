package dataManager

import (
	"errors"
	"reflect"

	"github.com/krix38/ScorchedGo/model/entity"
)

type Action int

type EntityCollection interface {
	GetCollection() []interface{}
}

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

type roomsCollection struct {
	roomsId int64
	rooms   entity.RoomsList
}

type playersCollection struct {
	playersId int64
	players   entity.PlayersList
}

func (r roomsCollection) GetCollection() []interface{} {
	entityCollection := make([]interface{}, len(r.rooms.Rooms)) // dlaczego nie zwrocic po prostu Rooms?
	for i := range r.rooms.Rooms {
		entityCollection[i] = r.rooms.Rooms[i]
	}
	return entityCollection
}

func (p playersCollection) GetCollection() []interface{} {
	entityCollection := make([]interface{}, len(p.players.Players))
	for i := range p.players.Players {
		entityCollection[i] = p.players.Players[i]
	}
	return entityCollection
}

var ActionChan = make(chan EntityAction)

var entityMemMapping = make(map[reflect.Type]EntityCollection)
var sharedPlayersCollection playersCollection
var sharedRoomsCollection roomsCollection

func init() {
	sharedPlayersCollection := playersCollection{
		players:   entity.PlayersList{Players: []entity.Player{}},
		playersId: 0,
	}
	sharedRoomsCollection := roomsCollection{
		rooms:   entity.RoomsList{Rooms: []entity.Room{}},
		roomsId: 0,
	}
	entityMemMapping[reflect.TypeOf(entity.Room{})] = sharedRoomsCollection
	entityMemMapping[reflect.TypeOf(entity.Player{})] = sharedPlayersCollection
}

func findEntity(id int64, entityType reflect.Type) (int, *interface{}) {
	for index, entity := range entityMemMapping[entityType].GetCollection() {
		if entity.Id == id {
			return index, &entity
		}
	}
	return -1, nil
}

func entityCreate(action EntityAction) {
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

func entityRead(action EntityAction) {
	switch action.AdditionalData.Mode {

	case ALL:
		action.ResponseChan <- memData.rooms
	case SINGLE:
		_, room := findRoom(action.AdditionalData.Index, memData)
		action.ResponseChan <- room
	}
}

func entityUpdate(action EntityAction) {
	room, ok := action.Entity.(entity.Room)
	if ok {
		index, _ := findRoom(room.Id, memData)
		memData.rooms.Rooms[index] = room
		action.ResponseChan <- nil
	} else {
		action.ResponseChan <- errors.New("failed to update room")
	}
}

func entityDelete(action EntityAction) {
	room, ok := action.Entity.(entity.Room)
	if ok {
		index, _ := findRoom(room.Id, memData)
		memData.rooms.Rooms = append(memData.rooms.Rooms[:index], memData.rooms.Rooms[index+1:]...)
		action.ResponseChan <- nil
	} else {
		action.ResponseChan <- errors.New("failed to delete room")
	}
}

func dispatchAction(action EntityAction) {
	switch action.Action {

	case CREATE:
		entityCreate(action)
	case READ:
		entityRead(action)
	case UPDATE:
		entityUpdate(action)
	case DELETE:
		entityDelete(action)
	}
}

func RunDataManager() {
	for {
		action := <-ActionChan
		dispatchAction(action)
	}
}
