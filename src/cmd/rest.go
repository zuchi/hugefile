package main

import (
	"90poe/src/pkg/infra/rest"
	"90poe/src/pkg/infra/startup"
)

func main() {
	dependencies := startup.InitDependencies()
	server := rest.NewServer(dependencies.PortUC)
	server.StartServer(":3000")
}
