package api

import (
	"github.com/gin-gonic/gin"

	"theaveasso.bab/internal/db"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store: store,
	}

	router := gin.Default()

    router.POST("/accounts", server.createAccount)
    router.GET("/accounts/:id", server.getAccount)
    router.GET("/accounts", server.listAccounts)
    router.POST("/users", server.createUser)
    router.GET("/users/:username", server.getUser)


	server.router = router
	return server
}

func (server *Server) Start(address string) error {
    return server.router.Run(address)
}

func errorResponse(err error) gin.H {
    return gin.H{"error": err.Error()}
}
