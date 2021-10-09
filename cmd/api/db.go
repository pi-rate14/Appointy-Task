package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	UsersCollection     *mongo.Collection
	PostsCollection   *mongo.Collection
	Ctx                 = context.TODO()
)

func CloseClientDB(client *mongo.Client ) {
    if client == nil {
        return
    }

    err := client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Connection to MongoDB closed.")
}

/*opens a database connection to mongodb*/
func openDB()  *mongo.Client {
	connectionURI := "mongodb+srv://pirate:pirate2546@cluster0.9u5da.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("Appointy")
	UsersCollection = db.Collection("Users")
	PostsCollection = db.Collection("Posts")
	return client

	
}
