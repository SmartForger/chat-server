package cmd

import (
	"fmt"
	"os"
)

func PrintAdminSecret() {
	publickey := os.Args[1]

	if publickey != "" {
		secret, err := EncryptRSA(GetAdminSecret(), publickey)

		if err == nil {
			fmt.Println(secret)
			return
		}
	}
	fmt.Println("Error")
}
