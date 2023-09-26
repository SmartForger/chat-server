package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupServer() {
	publ, priv, err := GenerateKey()
	if err != nil {
		panic(err)
	}

	CSet("m_public", publ)
	CSet("m_private", priv)
	CSet("m_secret", os.Getenv("CHATSERVER_SECRET"))

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	r.GET("/_key", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"key": CGet("m_public")})
	})

	r.POST("/client", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		secret := CreateClient(&user)
		c.JSON(http.StatusCreated, gin.H{"secret": secret})
	})

	r.GET("/client/:username", func(c *gin.Context) {
		username := c.Param("username")
		client := GetClient(username)

		c.JSON(http.StatusOK, gin.H{"client": client})
	})

	r.Run()
}
