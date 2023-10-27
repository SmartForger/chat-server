package cmd

import (
	"fmt"
	"os"

	lib "fullstackdevs14/chat-server/lib"

	"github.com/go-zoox/fetch"
)

func PrintAdminSecret() {
	serverUrl := lib.GetArg(os.Args, 3)
	if serverUrl == "" {
		serverUrl = "http://localhost:8080"
	}

	response, _ := fetch.Get(serverUrl + "/_key")

	keyJson := response.Value()

	publickey := keyJson.Get("key").String()

	if publickey != "" {
		secret, err := lib.EncryptRSA(lib.GetAdminSecret(), publickey)

		if err == nil {
			fmt.Println(secret)
			return
		}
	}
	fmt.Println("Error")
}

func PrintRSAKeys() {
	publ, priv, err := lib.GenerateKey()

	if err != nil {
		panic(err)
	}

	fmt.Println(publ)
	fmt.Println(priv)

	encrypted, _ := lib.EncryptRSA("Hello Test", publ)
	fmt.Println("Encrypted: " + encrypted)

	decrypted, _ := lib.DecryptRSA(encrypted, priv)
	fmt.Println("Decrypted: " + decrypted)
}
