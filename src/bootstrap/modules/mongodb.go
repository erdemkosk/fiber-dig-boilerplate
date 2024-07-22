package modules

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IMongoDB interface {
}

type MongoDB struct {
	Client *mongo.Client
}

func NewMongoDB(url string) (*MongoDB, error) {
	clientOptions := options.Client().ApplyURI(url)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	log.Println("MongoDB connected")

	return &MongoDB{Client: client}, nil
}
