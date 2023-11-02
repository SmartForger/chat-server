package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fullstackdevs14/chat-server/lib"
	"fullstackdevs14/chat-server/server/common"
)

func AdminReqMiddelware(c *gin.Context) {
	token := c.GetHeader(lib.HEADER_ADMIN_TOKEN)
	nonce := c.GetString(lib.HEADER_NONCE_TOKEN)

	if !lib.IsAdminToken(token, nonce) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{
			Message: "unauthorized",
		})
		return
	}

	c.Next()
}
