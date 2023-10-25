package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAdminRoutes(apiGroup *gin.RouterGroup) {
	adminGroup := apiGroup.Group("/admin")
	{
		adminGroup.Use(AdminReqMiddelware)

		adminGroup.POST("/client", func(c *gin.Context) {
			var user User
			if err := c.ShouldBindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			secret := CreateClient(&user)
			c.JSON(http.StatusCreated, gin.H{"secret": secret})
		})
	}
}
