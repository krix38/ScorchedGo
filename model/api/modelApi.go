package api

type ConnectionInfo struct {
	SignedIn bool
}

type Channel struct {
	Player1Id         int64
	Player2Id         int64
	ChannelName       string
	PasswordProtected bool
	Password          string
}

type ChannelsList struct {
	Channels []Channel
}
