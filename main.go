package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/abuzarsaddiqui/simplebank/api"
	db "github.com/abuzarsaddiqui/simplebank/db/sqlc"
	"github.com/abuzarsaddiqui/simplebank/gapi"
	"github.com/abuzarsaddiqui/simplebank/pb"
	"github.com/abuzarsaddiqui/simplebank/util"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	runGrpcServer(config, store)

}
func runGrpcServer(config util.Config, store *db.Store) {
	server, err := gapi.NewServer(config, *store)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener:", err)
	}
	log.Printf("start gRPC server at %s:", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}

}

func runGinServer(config util.Config, store *db.Store) {
	server, err := api.NewServer(config, *store)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
