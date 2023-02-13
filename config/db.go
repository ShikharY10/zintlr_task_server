package config

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataBase struct {
	Users *mongo.Collection
	Posts *mongo.Collection
}

func ConnectToDataBase(connStr string) *DataBase {

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(connStr).SetServerAPIOptions(serverAPIOptions)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	var db DataBase
	storage := client.Database("storage")
	db.Users = storage.Collection("users")
	db.Posts = storage.Collection("posts")

	return &db
}
