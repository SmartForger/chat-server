package common

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Client struct {
	Username string
	Password string
	Secret   string
}

type ErrorResponse struct {
	Message interface{} `json:"message"`
}

type SocketMessage struct {
	S string
	T string
}

type ChatMessage struct {
	Room    string
	Message string
}
