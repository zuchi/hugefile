package mongodb

import (
	"90poe/src/pkg/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PortRepositoryImpl struct {
	col *mongo.Collection
}

func NewPortRepositoryImpl(col *mongo.Collection) *PortRepositoryImpl {
	return &PortRepositoryImpl{col: col}
}

func (p PortRepositoryImpl) Save(ctx context.Context, port domain.Port) error {
	mongoPort := fromDomain(port)

	opts := options.Update().SetUpsert(true)
	_, err := p.col.UpdateOne(ctx, bson.M{"_id": port.Key}, bson.M{"$set": mongoPort}, opts)
	if err != nil {
		return err
	}
	return nil
}

func (p PortRepositoryImpl) FindByKey(ctx context.Context, key string) (domain.Port, error) {
	var mongoPort Port
	err := p.col.FindOne(ctx, bson.M{"_id": key}).Decode(&mongoPort)
	if err != nil && err != mongo.ErrNoDocuments {
		return domain.Port{}, err
	}

	return mongoPort.toDomain(), nil
}
