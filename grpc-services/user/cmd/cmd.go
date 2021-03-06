package cmd

import (
	commonCmd"github.com/omaressameldin/bet-collector-services/grpc-services/common/pkg/cmd"
	v1 "github.com/omaressameldin/bet-collector-services/grpc-services/user/pkg/service/v1"
	"github.com/omaressameldin/bet-collector-services/grpc-services/user/server"
)

var v1API *v1.UserServiceServer

// RunServer get the flags and starts the server
func RunServer() error {
	ctx, cfg, err := commonCmd.ConfigureFlags()
	if err != nil {
		return err
	}

	connector := commonCmd.InitFirebaseConnector(cfg.FirebaseConfig, cfg.Collection)
	v1API = v1.NewUserServiceServer(connector)

	return server.RunServer(ctx, v1API, cfg.Port)
}

// CloseServer closes all connections such as database connection
func CloseServer() error {
	return v1API.CloseConnection()
}
