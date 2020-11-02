package db

import (
	databaseconnectionmodel "bookstore/models/databaseconnection"
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect function establish the connection with mongodb returns a *databaseconnectionmodel.Connection instance or an error interface.
func Connect() (*databaseconnectionmodel.Connection, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DATABASE_URI")))

	if err != nil {
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)

	if err != nil {
		return nil, err
	}

	connection := databaseconnectionmodel.Connection{Client: client, Ctx: ctx}

	return &connection, nil
}
