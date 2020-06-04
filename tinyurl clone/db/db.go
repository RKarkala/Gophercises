package db

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
	"tinyurl-clone/models"
)

var col *mongo.Collection

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if col == nil {
		col = client.Database("url-db").Collection("maps")
	}
}

func Insert(data bson.M) bool {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := col.InsertOne(ctx, data)
	if err != nil {
		log.Fatal(err)
	}
	return true
}

func Find(query bson.D) (models.Entry, bool) {
	var result models.Entry
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := col.FindOne(ctx, query).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return result, false
		}
		log.Fatal(err)
	}
	return result, true
}
