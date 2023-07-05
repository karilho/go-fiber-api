package mongodb

import (
	"context"
	"github.com/karilho/go-fiber-api/src/configuration/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var (
	MONGODB_URL  = "MONGO_URL"
	MONGODB_NAME = "MONGO_DB_NAME"
)

func NewMongoConnection(ctx context.Context) (*mongo.Database, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(MONGODB_URL)))
	if err != nil {
		return nil, err
	}
	//ping é pra ve se está conectado ok.
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}
	logger.Info("Connected to MongoDB!")
	return client.Database(os.Getenv(MONGODB_NAME)), nil

}
