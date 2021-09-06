package main

import (
	"authService/config"
	"authService/protocol"
	"authService/repository"
	"authService/server"
	"authService/service"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main(){

	cfg := config.NewConfig()

	conn, err := sql.Open("postgres", fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DbPort,
		cfg.Host,
		cfg.User,
		cfg.Password,
		cfg.Dbname,
		cfg.Sslmode))
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

	protocol.RegisterAuthServiceServer(grpcServer,tokenHandler)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
