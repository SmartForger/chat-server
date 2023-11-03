package server

import (
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"fullstackdevs14/chat-server/lib"
	"fullstackdevs14/chat-server/server/admin"
	"fullstackdevs14/chat-server/server/client"
	"fullstackdevs14/chat-server/server/common"

	socketio "github.com/googollee/go-socket.io"
)

func Setup() {
	publ, priv, err := lib.GenerateKey()
	if err != nil {
		panic(err)
	}

	lib.CSet(lib.CK_PUBLIC, publ)
	lib.CSet(lib.CK_PRIVATE, priv)
	lib.CSet(lib.CK_ADMIN_SECRET, lib.GetAdminSecret())

	client.CreateClient(&common.User{
		Username: "john",
		Password: "test",
	})

	r := gin.Default()

	server := socketio.NewServer(nil)

	SocketHandlers(server)

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "sync_nonce", "authorization"}
	r.Use(cors.New(config))

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World")
	})

	r.GET("/_key", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"key": lib.CGet(lib.CK_PUBLIC)})
	})

	protectedRoutes := r.Group("/")
	{
		protectedRoutes.Use(common.NonceMiddleware)

		// Admin API Routes
		admin.AddAdminRoutes(protectedRoutes)

		// Client API Routes
		client.ClientApiRoutes(protectedRoutes)
	}

	r.StaticFS("/public", http.Dir("./frontend"))

	if err := r.Run(); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
