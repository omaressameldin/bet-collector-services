package cmd

import (
	"errors"
	"flag"
	"time"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	"github.com/omaressameldin/bet-collector-services/elasticsearch-service/app/server"
)

type Config struct {
	Port int
}

func configureFlags() (*Config, error) {
	cfg := &Config{}

	flag.IntVar(&cfg.Port, "port", 0, "port to bind to")
	flag.Parse()

	return cfg, nil
}

// RunServer get the flags and starts the server
func RunServer() error {
	cfg, err := configureFlags()
	if err != nil {
		return err
	}

	es, err := elasticsearch7.NewClient(elasticsearch7.Config{
		MaxRetries:           5,
		RetryOnStatus:        []int{500, 502, 503, 504},
		EnableRetryOnTimeout: true,
		RetryBackoff: func(i int) time.Duration {
			return time.Duration(i) * 3 * time.Second
		},
	})
	if err != nil {
		return err
	}

	info, err := es.Info()
	if err != nil {
		return err
	}

	if info.IsError() {
		return errors.New(info.String())
	}

	return server.RunServer(es, cfg.Port)
}
