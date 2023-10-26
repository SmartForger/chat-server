package lib

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func EncryptResponse(data interface{}, c *gin.Context) string {
	b, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	secret := c.GetString(HEADER_NONCE_TOKEN)

	return EncryptAESByte(b, secret)
}
