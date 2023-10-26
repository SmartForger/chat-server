package common

import (
	"bytes"
	"encoding/json"
	"fullstackdevs14/chat-server/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NonceMiddleware(c *gin.Context) {
	token := c.GetHeader(lib.HEADER_NONCE_TOKEN)

	decrypted, err := lib.DecryptRSA(token, lib.CGet(lib.CK_PRIVATE))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: "invalid nonce",
		})
		return
	}

	c.Set(lib.HEADER_NONCE_TOKEN, decrypted)
	c.Next()
}

func RequestBodyMiddleware(c *gin.Context) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	str := buf.String()

	if str == "" {
		c.Next()
		return
	}

	secret := c.GetString(lib.HEADER_NONCE_TOKEN)
	decrypted := lib.DecryptAES(str, secret)

	var data interface{}
	err := json.Unmarshal([]byte(decrypted), data)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, ErrorResponse{
			Message: "invalid data",
		})
		return
	}

	c.Set(lib.REQUEST_BODY, data)
	c.Next()
}
