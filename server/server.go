package server

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fullstackdevs14/chat-server/lib"
	"fullstackdevs14/chat-server/server/admin"
	"fullstackdevs14/chat-server/server/client"
	"fullstackdevs14/chat-server/server/common"
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

	protectedRoutes := r.Group("/")
	{
		protectedRoutes.Use(common.NonceMiddleware)
		protectedRoutes.Use(common.RequestBodyMiddleware)

		// Admin API Routes
		admin.AddAdminRoutes(protectedRoutes)

		// Client API Routes
		client.ClientApiRoutes(protectedRoutes)
	}

	r.Run()
}
