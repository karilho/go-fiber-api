package mongodb

import (
	"context"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitConnectionDB() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	//ping é pra ve se está conectado ok.
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	logger.Info("Connected to MongoDB!")
}
