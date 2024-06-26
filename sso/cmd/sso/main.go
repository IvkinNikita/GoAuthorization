package main

import (
	"fmt"
	"github.com/IvkinNikita/GoAuthorization/sso/sso/cmd/internal/config"
)

func main() {
	// TODO: initilization config object
	cfg := config.MustLoad()

	fmt.Println(cfg)
	// TODO: initilization logger

	// TODO: initilization apppp (app)

	// TODO: start gRPC-server app
}
