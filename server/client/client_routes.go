package client

import (
	"fullstackdevs14/chat-server/lib"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClientApiRoutes(apiGroup *gin.RouterGroup) {
	clientApiGroup := apiGroup.Group("/client")
	{
		clientApiGroup.GET("/:username", func(c *gin.Context) {
			username := c.Param("username")
			client := GetClient(username)

			c.JSON(http.StatusOK, lib.EncryptResponse(client, c))
		})
	}
}
