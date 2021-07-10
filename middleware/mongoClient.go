package middleware

import (
	"context"

	"github.com/labstack/echo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	ClientKey        = "MongoClient"
	ClientContextKey = "MongoContex"
)

func MongoClientHandler(client *mongo.Client, ctx context.Context) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {

			c.Set(ClientKey, client)
			c.Set(ClientContextKey, ctx)

			err := client.Ping(ctx, readpref.Primary())
			if err != nil {
				return err
			}

			return nil
		})
	}
}
