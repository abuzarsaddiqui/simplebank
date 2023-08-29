package gapi

import (
	"fmt"

	db "github.com/abuzarsaddiqui/simplebank/db/sqlc"
	"github.com/abuzarsaddiqui/simplebank/pb"
	"github.com/abuzarsaddiqui/simplebank/token"
	"github.com/abuzarsaddiqui/simplebank/util"
)

// Server serves gRPC requests
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

func NewServer(config util.Config, str db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker %w", err)
	}

	server := &Server{store: str, tokenMaker: tokenMaker, config: config}

	return server, nil
}
