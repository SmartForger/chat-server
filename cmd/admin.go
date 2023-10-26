package cmd

import (
	"fmt"
	"os"

	lib "fullstackdevs14/chat-server/lib"
)

func PrintAdminSecret() {
	publickey := os.Args[2]

	if publickey != "" {
		secret, err := lib.EncryptRSA(lib.GetAdminSecret(), publickey)

		if err == nil {
			fmt.Println(secret)
			return
		}
	}
	fmt.Println("Error")
}
