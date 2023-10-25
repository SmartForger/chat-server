package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupServer() {
	publ, priv, err := GenerateKey()
	if err != nil {
		panic(err)
	}

	CSet(CK_PUBLIC, publ)
	CSet(CK_PRIVATE, priv)
	CSet(CK_ADMIN_SECRET, GetAdminSecret())

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	r.GET("/_key", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"key": CGet(CK_PUBLIC)})
	})

	r.GET("/client/:username", func(c *gin.Context) {
		username := c.Param("username")
		client := GetClient(username)

		c.JSON(http.StatusOK, gin.H{"client": client})
	})

	AddAdminRoutes(&r.RouterGroup)

	r.Run()
}
