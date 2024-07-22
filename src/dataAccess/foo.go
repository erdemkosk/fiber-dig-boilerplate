package dataAccess

import (
	"context"
	"fiber-boilerplate/src/bootstrap/modules"

	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/dig"
)

type IFooDataAccess interface {
	GetFoo(id string) (string, error)
}

type FooDataAccess struct {
	mongoDB *modules.MongoDB
	//env     *config.Env // Add config.Env to reach config
}

type FooDataAccessDependencies struct {
	dig.In

	MongoDB *modules.MongoDB
	//Env     *config.Env
}

func NewFooDataAccess(deps FooDataAccessDependencies) *FooDataAccess {
	return &FooDataAccess{
		mongoDB: deps.MongoDB,
		//env:     deps.Env,
	}
}

func (dataAccess *FooDataAccess) GetFoo(id string) (string, error) {
	collection := dataAccess.mongoDB.Client.Database("fiber").Collection("foo")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("Invalid ID format: %v", err)
		return "", err
	}

	filter := bson.M{"_id": objectID}

	var result struct {
		Name string `bson:"name"`
	}

	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return "", nil
		}
		log.Printf("Failed to find document: %v", err)
		return "", err
	}

	return result.Name, nil
}
