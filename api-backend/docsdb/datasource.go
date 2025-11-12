package docsdb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDataSource(uri string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	opts.SetMaxPoolSize(100)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return nil, err
	}
	return client, nil

}
