package cmd

import (
	v1 "github.com/omaressameldin/bet-collector-services/grpc-services/bet/pkg/service/v1"
	"github.com/omaressameldin/bet-collector-services/grpc-services/bet/server"
	commonCmd "github.com/omaressameldin/bet-collector-services/grpc-services/common/pkg/cmd"
)

var v1API *v1.BetServiceServer

// RunServer get the flags and starts the server
func RunServer() error {
	ctx, cfg, err := commonCmd.ConfigureFlags()
	if err != nil {
		return err
	}

	connector := commonCmd.InitFirebaseConnector(cfg.FirebaseConfig, cfg.Collection)
	v1API = v1.NewBetServiceServer(connector, cfg.Dependencies)

	return server.RunServer(ctx, v1API, cfg.Port)
}

// CloseServer closes all connections such as database connection
func CloseServer() error {
	return v1API.CloseConnection()
}
