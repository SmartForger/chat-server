package client

import (
	"fullstackdevs14/chat-server/server/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminApiRoutes(apiGroup *gin.RouterGroup) {
	clientAdminApiGroup := apiGroup.Group("/client")
	{
		clientAdminApiGroup.POST("/", func(c *gin.Context) {
			var user common.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			secret := CreateClient(&user)
			c.JSON(http.StatusCreated, gin.H{"secret": secret})
		})
	}
}
