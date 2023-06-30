package main

import (
	"90poe/src/pkg/infra/rest"
	"90poe/src/pkg/infra/startup"
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	log := logger.Sugar()

	ctx := context.Background()
	dependencies := startup.InitDependencies(ctx)
	server := rest.NewServer(dependencies.PortUC)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt, os.Kill)

	go server.Run(":3000")

	<-quit

	log.Infof("shutdown in progress...")
	newContext, timeoutFunc := context.WithTimeout(ctx, 5*time.Second)

	log.Infof("closing api server...")
	err := server.Shutdown(newContext)
	if err != nil {
		log.Errorf("cannot shutdown api server: %v", err)
	}

	log.Infof("closing mongo server...")
	err = dependencies.MongoClient.Disconnect(newContext)
	if err != nil {
		log.Errorf("cannot shutdown mongo server: %v", err)
	}

	defer timeoutFunc()
	log.Infof("shutdown finished")

}
