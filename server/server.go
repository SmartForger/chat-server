package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fullstackdevs14/chat-server/lib"
	"fullstackdevs14/chat-server/server/admin"
	"fullstackdevs14/chat-server/server/client"
)

func Setup() {
	publ, priv, err := lib.GenerateKey()
	if err != nil {
		panic(err)
	}

	lib.CSet(lib.CK_PUBLIC, publ)
	lib.CSet(lib.CK_PRIVATE, priv)
	lib.CSet(lib.CK_ADMIN_SECRET, lib.GetAdminSecret())

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	r.GET("/_key", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"key": lib.CGet(lib.CK_PUBLIC)})
	})

	r.GET("/client/:username", func(c *gin.Context) {
		username := c.Param("username")
		client := client.GetClient(username)

		c.JSON(http.StatusOK, gin.H{"client": client})
	})

	admin.AddAdminRoutes(&r.RouterGroup)

	r.Run()
}
