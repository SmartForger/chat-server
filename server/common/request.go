package common

import (
	"bytes"
	"encoding/json"
	"fullstackdevs14/chat-server/lib"

	"github.com/gin-gonic/gin"
)

func GetRequestBody[T interface{}](c *gin.Context) (T, bool) {
	var data T

	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	str := buf.String()

	if str == "" {
		return data, false
	}

	secret := c.GetString(lib.HEADER_NONCE_TOKEN)
	decrypted := lib.DecryptAES(str, secret)

	err := json.Unmarshal([]byte(decrypted), &data)

	if err != nil {
		return data, false
	}

	return data, true
}
