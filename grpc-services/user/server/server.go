package server

import (
	"context"

	commonServer "github.com/omaressameldin/bet-collector-services/grpc-services/common/pkg/server"
	v1 "github.com/omaressameldin/bet-collector-services/grpc-services/user/pkg/api/v1"
	"google.golang.org/grpc"
)

// RunServer runs service to publish User service
func RunServer(ctx context.Context, v1API v1.UserServiceServer, port string) error {
	server := grpc.NewServer()
	v1.RegisterUserServiceServer(server, v1API)

	return commonServer.RunServer(ctx, server, port)
}
