package dao

import (
	"github.com/krix38/ScorchedGo/model/entity"
)

func createChannel(hostPlayer int64, channelName string, password string) {
	var newChannel entity.Channel
	if password == "" {
		newChannel = entity.Channel{
			Player1Id:         hostPlayer,
			ChannelName:       channelName,
			PasswordProtected: false,
		}
	} else {
		newChannel = entity.Channel{
			Player1Id:         hostPlayer,
			ChannelName:       channelName,
			PasswordProtected: true,
			Password:          password,
		}
	}
	
}
