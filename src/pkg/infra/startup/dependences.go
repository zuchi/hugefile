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
	"os"
	"time"
)

type Dependencies struct {
	PortParser  ports.PortParser
	PortUC      *use_cases.PortUC
	MongoClient *mongo.Client
}

func InitDependencies(ctx context.Context) Dependencies {
	ctxTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	mongoURI := os.Getenv("DATABASE_URI")

	mongoClient, err := mongo.Connect(ctxTimeout, options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(fmt.Errorf("cannot connect to the database: %w", err))
	}

	err = mongoClient.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Errorf("cannot ping into mongodb database: %w", err))
	}

	databaseName := os.Getenv("DATABASE_NAME")
	portCollection := mongoClient.Database(databaseName).Collection("port")
	portRepositoryImpl := mongodb.NewPortRepositoryImpl(portCollection)

	portService := ports.NewServicePort(portRepositoryImpl)

	dep := Dependencies{}
	dep.PortParser = port_parser.NewJsonParser()
	dep.PortUC = use_cases.NewPortUC(dep.PortParser, portService)
	dep.MongoClient = mongoClient
	return dep
}
