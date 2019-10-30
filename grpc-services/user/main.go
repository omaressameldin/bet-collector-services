package main

import (
	"fmt"
	"os"

	"github.com/omaressameldin/bet-collector-services/grpc-services/user/cmd"
)

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func close() {
	if err := cmd.CloseServer(); err != nil {
		exitWithError(err)
	}
}

func main() {
	defer close()

	if err := cmd.RunServer(); err != nil {
		exitWithError(err)
	}
}
