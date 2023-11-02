package common

import (
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
