package main

import (
	"database/sql"
	"fmt"
	"github.com/INEFFABLE-games/authService/internal/config"
	"github.com/INEFFABLE-games/authService/internal/repository"
	"github.com/INEFFABLE-games/authService/internal/server"
	"github.com/INEFFABLE-games/authService/internal/service"
	"github.com/INEFFABLE-games/authService/protocol"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {

	cfg := config.NewConfig()

	conn, err := sql.Open("postgres", cfg.SqlConnString)
	if err != nil {
		log.Errorf("main: unable to open sql connection %v,", err)
	}

	if err := conn.Ping(); err != nil {
		log.Errorf("main: unable to ping sql connection %v,", err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", cfg.GrpcPort))
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "main",
			"action":  "create listener",
		}).Errorf("unable to create listener %v", err.Error())
	}

	tokenRepo := repository.NewTokenRepository(conn)
	tokenService := service.NewAuthService(tokenRepo)
	tokenHandler := server.NewAuthServer(tokenService)

	grpcServer := grpc.NewServer()

	protocol.RegisterAuthServiceServer(grpcServer, tokenHandler)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
