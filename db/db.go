package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(url string) (*mongo.Client, *context.Context) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(url))
	return client, &ctx
}

func DisconnectDB(client *mongo.Client, ctx *context.Context) {
	client.Disconnect(*ctx)
}
