package main

import (
	"github.com/LoomingLunar/LunarLoom-chat-service/connection"
	"github.com/joho/godotenv"
)

func main() {
	// Loading enviornment variables
	godotenv.Load(".env")

	// Connecting database
	connection.RedisSetUp()
}
