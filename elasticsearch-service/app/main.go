package main

import (
	"fmt"
	"os"

	"github.com/omaressameldin/bet-collector-services/elasticsearch-service/app/cmd"
)

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func main() {
	if err := cmd.RunServer(); err != nil {
		exitWithError(err)
	}
}
