// utils/mongodb.go
package utils

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectMongoDB initializes a connection to MongoDB using the provided URI
func ConnectMongoDB(uri string) {
    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.NewClient(clientOptions)
    if err != nil {
        log.Fatalf("Failed to create MongoDB client: %v", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Ping the database to verify the connection
    if err := client.Ping(ctx, nil); err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    log.Println("Connected to MongoDB!")
    Client = client
}
