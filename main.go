package main

import (
	"os"

	"fullstackdevs14/chat-server/cmd"
	"fullstackdevs14/chat-server/lib"
	"fullstackdevs14/chat-server/server"
)

func main() {
	command := lib.GetArg(os.Args, 1)

	switch command {
	case "cli":
		cmd.Run()
	default:
		server.Setup()
	}
}
