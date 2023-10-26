package client

import (
	"fmt"

	"fullstackdevs14/chat-server/lib"
	"fullstackdevs14/chat-server/server/common"
)

func CreateClient(user *common.User) string {
	lib.CSet(fmt.Sprintf("c:%s:u", user.Username), user.Username)
	lib.CSet(fmt.Sprintf("c:%s:p", user.Username), user.Password)

	secret := lib.GenerateAESKey()
	lib.CSet(fmt.Sprintf("c:%s:k", user.Username), secret)

	return secret
}

func GetClient(username string) *common.Client {
	password := lib.CGet(fmt.Sprintf("c:%s:p", username))
	secret := lib.CGet(fmt.Sprintf("c:%s:k", username))

	return &common.Client{Username: username, Password: password, Secret: secret}
}
