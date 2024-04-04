package handlers

import (
	databasehelpers "github.com/Qryptell/Qryptell-chat-service/internal/helpers/databaseHelpers"
	"github.com/Qryptell/Qryptell-chat-service/pkg/message"
	"github.com/Qryptell/Qryptell-chat-service/pkg/redis"
)

// Handling messages send by the user
func HandleUserMessages() {
  // Channel for messages
  var userMessages chan message.ServerMsg = make(chan message.ServerMsg)

  // Listening to messages in redis
  go redis.ReadMessages(redis.RECEIVE_USER_CHANNEL, userMessages)

	// Handling messages
  for {
    select {
    case msg := <-userMessages:
			databasehelpers.InsertTextMessage(msg.Msg)
    }
  }
}
