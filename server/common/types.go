package common

type User struct {
	USERNAME string `json:"username" binding:"required"`
	PASSWORD string `json:"password" binding:"required"`
}

type Client struct {
	USERNAME string
	PASSWORD string
	SECRET   string
}

type ErrorResponse struct {
	Message interface{} `json:"message"`
}