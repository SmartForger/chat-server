package main

import (
	"os"

	"fullstackdevs14/chat-server/cmd"
	"fullstackdevs14/chat-server/server"
)

func main() {
	args := os.Args

	var command string
	if len(args) > 1 {
		command = args[1]
	} else {
		command = ""
	}

	switch command {
	case "cli":
		cmd.Run()
	default:
		server.Setup()
	}
}
