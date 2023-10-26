package admin

import (
	"github.com/gin-gonic/gin"

	"fullstackdevs14/chat-server/server/client"
)

func AddAdminRoutes(apiGroup *gin.RouterGroup) {
	adminGroup := apiGroup.Group("/admin")
	{
		adminGroup.Use(AdminReqMiddelware)

		client.AdminApiRoutes(adminGroup)
	}
}
