package client

import (
	"fmt"

	"fullstackdevs14/chat-server/lib"
	"fullstackdevs14/chat-server/server/common"
)

func CreateClient(user *common.User) string {
	lib.CSet(fmt.Sprintf("c:%s:u", user.USERNAME), user.USERNAME)
	lib.CSet(fmt.Sprintf("c:%s:p", user.USERNAME), user.PASSWORD)

	secret := lib.GenerateAESKey()
	lib.CSet(fmt.Sprintf("c:%s:k", user.USERNAME), secret)

	return secret
}

func GetClient(username string) *common.Client {
	password := lib.CGet(fmt.Sprintf("c:%s:p", username))
	secret := lib.CGet(fmt.Sprintf("c:%s:k", username))

	return &common.Client{USERNAME: username, PASSWORD: password, SECRET: secret}
}
