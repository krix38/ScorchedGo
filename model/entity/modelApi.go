package entity

type ConnectionInfo struct {
	SignedIn bool
}

type Player struct {
	Id   int64
	Name string
}

type PlayersList struct {
	Players []Player
}

type Room struct {
	Id                int64
	Player1Id         int64
	Player2Id         int64
	ChannelName       string
	PasswordProtected bool
	Password          string
}

type RoomsList struct {
	Rooms []Room
}
