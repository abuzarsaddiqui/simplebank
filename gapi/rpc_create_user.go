package gapi

import (
	"context"

	db "github.com/abuzarsaddiqui/simplebank/db/sqlc"
	"github.com/abuzarsaddiqui/simplebank/pb"
	"github.com/abuzarsaddiqui/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	password, err := util.HashedPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error")
	}
	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: password,
		Email:          req.Email,
		FullName:       req.FullName,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "internal error")
	}
	return &pb.CreateUserResponse{
		User: util.ConvertUser(user),
	}, nil
}
