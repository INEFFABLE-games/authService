package server

import (
	proto "authService/protocol"
	"authService/service"
	"context"
	log "github.com/sirupsen/logrus"
)

type AuthServer struct {
	authService *service.AuthService

	proto.UnimplementedAuthServiceServer
}

func (a AuthServer) Refresh(ctx context.Context, request *proto.RefreshRequest) (*proto.RefreshReply, error) {

	login := request.GetLogin()

	//----------------------------------------------------------------
	newJWT, newRT, err := a.authService.RefreshTokens(ctx, login)
	if err != nil {
		log.WithFields(log.Fields{
			"handler": "server",
			"action":  "refresh",
		}).Errorf("unable to create tokens %v", err.Error())

		newJWT := ""
		newRT := ""

		return &proto.RefreshReply{
			Jwt: &newJWT,
			Rt:  &newRT,
		}, err
	}

	return &proto.RefreshReply{
		Jwt: &newJWT,
		Rt:  &newRT,
	}, err
}

func NewAuthServer(authService *service.AuthService) proto.AuthServiceServer {
	return AuthServer{authService: authService}
}