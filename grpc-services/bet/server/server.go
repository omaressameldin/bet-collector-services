package server

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	v1 "github.com/omaressameldin/bet-collector-services/grpc-services/bet/pkg/api/v1"
	"google.golang.org/grpc"
)

// RunServer runs service to publish Bet service
func RunServer(ctx context.Context, v1API v1.BetServiceServer, port string) error {
	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	v1.RegisterBetServiceServer(server, v1API)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down bet server...")

			server.GracefulStop()

			<-ctx.Done()
		}
	}()

	log.Println("starting bet server...")
	return server.Serve(listen)
}
