package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Collection *mongo.Collection

func Connect() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(("mongodb://127.0.0.1:27017")))
	if err != nil {
		return err
	}

	Collection = client.Database("firstDB").Collection("users")

	return nil
}
