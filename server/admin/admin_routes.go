package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fullstackdevs14/chat-server/server/client"
	"fullstackdevs14/chat-server/server/common"
)

func AddAdminRoutes(apiGroup *gin.RouterGroup) {
	adminGroup := apiGroup.Group("/admin")
	{
		adminGroup.Use(AdminReqMiddelware)

		adminGroup.POST("/client", func(c *gin.Context) {
			var user common.User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			secret := client.CreateClient(&user)
			c.JSON(http.StatusCreated, gin.H{"secret": secret})
		})
	}
}
