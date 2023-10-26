package client

import (
	"fullstackdevs14/chat-server/lib"
	"fullstackdevs14/chat-server/server/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminApiRoutes(apiGroup *gin.RouterGroup) {
	clientAdminApiGroup := apiGroup.Group("/client")
	{
		clientAdminApiGroup.POST("/", func(c *gin.Context) {
			user, ok := common.GetRequestBody[common.User](c)

			if !ok {
				c.AbortWithStatusJSON(http.StatusBadRequest, common.ErrorResponse{
					Message: "invalid data",
				})
				return
			}

			secret := CreateClient(&user)
			c.JSON(http.StatusCreated, lib.EncryptResponse(gin.H{"secret": secret}, c))
		})
	}
}
