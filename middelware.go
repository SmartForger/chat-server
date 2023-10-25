package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminReqMiddelware(c *gin.Context) {
	token := c.GetHeader(HEADER_ADMIN_TOKEN)

	if !IsAdminToken(token) {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Message: "unauthorized",
		})
		return
	}

	c.Next()
}
