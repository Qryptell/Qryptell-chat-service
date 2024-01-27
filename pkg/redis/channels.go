package redis

// Type channel
type channel string

// Channel names
const RECEIVE_USER_CHANNEL channel = "USER_MESSAGE"

// Listening to redis channels
func ListenChannels() {
	go ReadMessages(RECEIVE_USER_CHANNEL)
}
