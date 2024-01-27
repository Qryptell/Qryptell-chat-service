/*
				Copyright © 2023 LunarLoom 
LunarLoom Web Socket Service - WebSocket Service for LunarLoom End To End Encrypted Chat App

*/
package main

import (
	"github.com/LoomingLunar/LunarLoom-chat-service/pkg/redis"
	"github.com/joho/godotenv"
)

func main() {
	// Loading enviornment variables
	godotenv.Load(".env")

	// Connecting redis
	redis.SetUp()

	// Keep server running
	var ch = make(chan bool)
	<-ch
}
