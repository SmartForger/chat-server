package server

import (
	"encoding/json"
	"fmt"
	"fullstackdevs14/chat-server/server/client"
	"fullstackdevs14/chat-server/server/common"

	socketio "github.com/googollee/go-socket.io"
)

func SocketHandlers(server *socketio.Server) {
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "join", func(s socketio.Conn, msg string) error {
		data, ok := common.GetSocketMessage[common.User](msg)

		if ok {
			s.Join(data.Username)
			server.BroadcastToRoom("/", data.Username, "joined", data.Username)
		} else {
			s.Emit("custom_error", "join_error")
		}

		return nil
	})

	server.OnEvent("/", "broadcast", func(s socketio.Conn, msg string) error {
		data, ok := common.GetSocketMessage[common.ChatMessage](msg)

		fmt.Println(data)

		if ok {
			json, err := json.Marshal(data)
			if err != nil {
				return nil
			}

			c := client.GetClient(data.Room)
			if c.Secret == "" {
				return nil
			}

			encrypted, ok := common.GetEncryptedData(json, c.Secret)
			if !ok {
				return nil
			}

			server.BroadcastToRoom("/", data.Room, "receive", encrypted)
		}

		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("closed", msg)
	})
}
