package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/omaressameldin/go-database-connector/app/pkg/database"
	"github.com/omaressameldin/go-database-connector/app/pkg/firebase"
)

type Config struct {
	Port           string
	FirebaseConfig string
	Collection     string
}

func ConfigureFlags() (context.Context, *Config, error) {
	ctx := context.Background()

	var cfg Config
	flag.StringVar(&cfg.Port, "port", "", "port to bind")
	flag.StringVar(&cfg.FirebaseConfig, "firebaseConfig", "", "firebase json config file")
	flag.StringVar(&cfg.Collection, "collection", "", "firebase collection")
	flag.Parse()

	if len(cfg.Port) == 0 {
		return nil, nil, fmt.Errorf("invalid TCP port for server: '%s'", cfg.Port)
	}

	if len(cfg.Collection) == 0 {
		return nil, nil, fmt.Errorf("invalid Collection for firebase database: '%s'", cfg.Collection)
	}

	_, err := os.Stat(cfg.FirebaseConfig)
	if os.IsNotExist(err) {
		return nil, nil, fmt.Errorf("File does not exist: '%s'", cfg.FirebaseConfig)
	}

	return ctx, &cfg, nil
}

func InitFirebaseConnector(firebaseConfig, collection string) database.Connector {
	connector, err := firebase.StartConnection(firebaseConfig, collection, "")
	if err != nil {
		panic(err)
	}

	return connector
}
