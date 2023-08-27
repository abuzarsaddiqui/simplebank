package api

import (
	db "github.com/abuzarsaddiqui/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(str *db.Store) *Server {
	server := &Server{store: str}
	router := gin.Default()
	// add routes to router
	router.POST("/account", server.CreateAccount)
	router.GET("/account/:id", server.GetAccount)
	router.GET("/account", server.ListAccounts)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
