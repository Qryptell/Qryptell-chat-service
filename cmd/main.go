/*
	Copyright © 2023 LunarLoom

LunarLoom Chat Service - Chat Service for LunarLoom End To End Encrypted Chat App
*/
package main

import (
	"github.com/LoomingLunar/LunarLoom-chat-service/internal/handlers"
	"github.com/LoomingLunar/LunarLoom-chat-service/pkg/redis"
	"github.com/LoomingLunar/LunarLoom-chat-service/pkg/scylla"
	"github.com/joho/godotenv"
)

func main() {
	// Loading enviornment variables
	godotenv.Load(".env")

	// Connecting database and redis
	redis.SetUp()
	scylla.SetUp()

	handlers.HandleUserMessages()
}
