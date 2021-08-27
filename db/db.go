package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(url string) (*mongo.Client, *context.Context) {
	ctx := context.Background()
	client, err := mongo.NewClient(options.Client().ApplyURI(url))

	if err != nil {
		fmt.Println("mongo client create error:", err)
	}

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println("mongo client connection error:", err)
	}

	return client, &ctx
}

func DisconnectDB(client *mongo.Client, ctx *context.Context) {
	client.Disconnect(*ctx)
}
