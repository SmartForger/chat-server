package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupServer() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	r.Run()
}
