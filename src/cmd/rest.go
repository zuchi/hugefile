package main

import (
	"90poe/src/pkg/infra/rest"
	"90poe/src/pkg/infra/startup"
	"context"
)

func main() {
	ctx := context.Background()
	dependencies := startup.InitDependencies(ctx)
	server := rest.NewServer(dependencies.PortUC)
	server.StartServer(":3000")
}
