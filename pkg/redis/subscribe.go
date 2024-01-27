package redis

import (
	"context"
	"encoding/json"

	"github.com/LoomingLunar/LunarLoom-chat-service/pkg/message"
)

// Read messages from channel
func ReadMessages(ch channel) {
	// Subscribing channel
	var pubsub = Redis.Subscribe(context.TODO(), string(ch))
	defer pubsub.Close()

	// Listening to messages to the channel
	var c = pubsub.Channel()
	for m := range c {
		var msg message.ServerMsg
		json.Unmarshal([]byte(m.Payload), &msg)
	}
}
