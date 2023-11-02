package client

import (
	"fullstackdevs14/chat-server/lib"
	"fullstackdevs14/chat-server/server/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClientApiRoutes(apiGroup *gin.RouterGroup) {
	clientApiGroup := apiGroup.Group("/client")
	{
		clientApiGroup.POST("/login", func(c *gin.Context) {
			user, ok := common.GetRequestBody[common.User](c)

			if !ok {
				c.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponse{
					Message: "invalid data",
				})
				return
			}

			client := GetClient(user.Username)

			if user.Password != client.Password {
				c.JSON(http.StatusUnauthorized, common.ErrorResponse{
					Message: "unauthorized",
				})
				return
			}

			c.JSON(http.StatusOK, lib.EncryptResponse(client, c))
		})
	}
}
