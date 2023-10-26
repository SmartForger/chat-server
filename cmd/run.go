package cmd

import (
	"fullstackdevs14/chat-server/lib"
	"os"
)

func Run() {
	command := lib.GetArg(os.Args, 2)

	switch command {
	case "adminsecret":
		PrintAdminSecret()
	default:
	}
}
