/*
	Copyright © 2023 Qryptell

Qryptell Chat Service - Chat Service for Qryptell End To End Encrypted Chat App
*/
package main

import (
	"github.com/Qryptell/Qryptell-chat-service/internal/handlers"
	"github.com/Qryptell/Qryptell-chat-service/pkg/redis"
	"github.com/Qryptell/Qryptell-chat-service/pkg/scylla"
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
