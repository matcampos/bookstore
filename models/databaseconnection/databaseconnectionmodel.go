package databaseconnectionmodel

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type Connection struct {
	Client *mongo.Client
	Ctx    context.Context
}
