package api

import (
	"fmt"

	db "github.com/abuzarsaddiqui/simplebank/db/sqlc"
	"github.com/abuzarsaddiqui/simplebank/token"
	"github.com/abuzarsaddiqui/simplebank/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

func NewServer(config util.Config, str db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	server := &Server{store: str, tokenMaker: tokenMaker, config: config}

	//add validtors
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}
	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	// add routes to router
	router.POST("/users", server.CreateUser)
	router.POST("/login", server.LoginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.POST("/account", server.CreateAccount)
	authRoutes.GET("/account/:id", server.GetAccount)
	authRoutes.GET("/account", server.ListAccounts)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
