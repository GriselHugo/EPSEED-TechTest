package main

import (
	"main/internal/config"
	"main/internal/server"
)

func main() {
	config.LoadEnv()
	server.InitializeServer()
}