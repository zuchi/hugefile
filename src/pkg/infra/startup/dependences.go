package startup

import (
	"90poe/src/pkg/domain/ports"
	"90poe/src/pkg/domain/use_cases"
	"90poe/src/pkg/infra/parsers/port_parser"
	"90poe/src/pkg/infra/repository/mongodb"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Dependencies struct {
	PortParser ports.PortParser
	PortUC     *use_cases.PortUC
}

func InitDependencies(ctx context.Context) Dependencies {
	ctxTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	mongoClient, err := mongo.Connect(ctxTimeout, options.Client().ApplyURI("mongodb://localhost:27017/port_db"))
	if err != nil {
		panic(fmt.Errorf("cannot connect to the database: %w", err))
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Errorf("cannot ping into mongodb database: %w", err))
	}

	portCollection := mongoClient.Database("port_db").Collection("port")
	portRepositoryImpl := mongodb.NewPortRepositoryImpl(portCollection)

	portService := ports.NewServicePort(portRepositoryImpl)

	dep := Dependencies{}
	dep.PortParser = port_parser.NewJsonParser()
	dep.PortUC = use_cases.NewPortUC(dep.PortParser, portService)
	return dep
}
