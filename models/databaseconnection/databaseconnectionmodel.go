package databaseconnectionmodel

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Connection receives the parameter Client which is the Client connection with mongodb and the Ctx is the context instance that the connection was associated.
type Connection struct {
	Client *mongo.Client
	Ctx    context.Context
}
