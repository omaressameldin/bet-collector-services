package server

import (
	"context"

	v1 "github.com/omaressameldin/bet-collector-services/grpc-services/bet/pkg/api/v1"
	commonServer "github.com/omaressameldin/bet-collector-services/grpc-services/common/pkg/server"
	"google.golang.org/grpc"
)

// RunServer runs service to publish Bet service
func RunServer(ctx context.Context, v1API v1.BetServiceServer, port string) error {
	server := grpc.NewServer()
	v1.RegisterBetServiceServer(server, v1API)

	return commonServer.RunServer(ctx, server, port)
}
