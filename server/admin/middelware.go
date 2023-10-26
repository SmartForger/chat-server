package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fullstackdevs14/chat-server/lib"
	"fullstackdevs14/chat-server/server/common"
)

func AdminReqMiddelware(c *gin.Context) {
	token := c.GetHeader(lib.HEADER_ADMIN_TOKEN)

	if !lib.IsAdminToken(token) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, common.ErrorResponse{
			Message: "unauthorized",
		})
		return
	}

	c.Next()
}
