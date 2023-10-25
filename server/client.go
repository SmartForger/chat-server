package main

import "fmt"

func CreateClient(user *User) string {
	CSet(fmt.Sprintf("c:%s:u", user.USERNAME), user.USERNAME)
	CSet(fmt.Sprintf("c:%s:p", user.USERNAME), user.PASSWORD)

	secret := GenerateAESKey()
	CSet(fmt.Sprintf("c:%s:k", user.USERNAME), secret)

	return secret
}

func GetClient(username string) *Client {
	password := CGet(fmt.Sprintf("c:%s:p", username))
	secret := CGet(fmt.Sprintf("c:%s:k", username))

	return &Client{username, password, secret}
}
