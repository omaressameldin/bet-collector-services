package main

import (
	"github.com/omaressameldin/bet-collector-services/grpc-services/common/pkg/server"
	"github.com/omaressameldin/bet-collector-services/grpc-services/user/cmd"
)

func close() {
	if err := cmd.CloseServer(); err != nil {
		server.ExitWithError(err)
	}
}

func main() {
	defer close()

	if err := cmd.RunServer(); err != nil {
		server.ExitWithError(err)
	}
}
