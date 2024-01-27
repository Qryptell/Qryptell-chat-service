package handlers

import (
	"fmt"

	"github.com/LoomingLunar/LunarLoom-chat-service/pkg/message"
	"github.com/LoomingLunar/LunarLoom-chat-service/pkg/redis"
)

func HandleUserMessages() {
	// Channel for messages
	var userMessages chan message.ServerMsg = make(chan message.ServerMsg)

	// Listening to messages in redis
	redis.ReadMessages(redis.RECEIVE_USER_CHANNEL, userMessages)

	for {
		select {
		case msg := <-userMessages:
			fmt.Println(msg)

		}
	}
}
