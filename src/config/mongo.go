package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDBClient *mongo.Client

// Kết nối đến MongoDB
func ConnectMongoDB(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Không thể kết nối MongoDB: %v", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Không thể ping MongoDB: %v", err)
	}

	fmt.Println("Đã kết nối MongoDB!")
	MongoDBClient = client
}

// Lấy collection từ MongoDB
func GetCollection(database, collection string) *mongo.Collection {
	return MongoDBClient.Database(database).Collection(collection)
}
