package redis

import (
	"context"
	"encoding/json"

	"github.com/LoomingLunar/LunarLoom-chat-service/pkg/message"
)

// Read messages from channel
func ReadMessages(channelName channel,ch chan <- message.ServerMsg) {
	// Subscribing channel
	var pubsub = Redis.Subscribe(context.TODO(), string(channelName))
	defer pubsub.Close()

	// Listening to messages to the channel
	var c = pubsub.Channel()
	for m := range c {
		var msg message.ServerMsg
		json.Unmarshal([]byte(m.Payload), &msg)
		ch <- msg
	}
}
