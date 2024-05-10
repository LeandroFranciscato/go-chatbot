package repo

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func New() {

	client, err := mongo.Connect(
		context.Background(),
		options.Client().
			ApplyURI("mongodb://localhost:27017").
			SetAuth(options.Credential{
				Username: "root",
				Password: "example",
			}),
	)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	db := client.Database("chatbot")
	collection := db.Collection("customer")
	cur, err := collection.Find(ctx, bson.D{})
	defer func() {
		_ = cur.Close(ctx)
	}()
	if err != nil {
		panic(err)
	}
	for cur.Next(context.Background()) {
		var result bson.D
		err := cur.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("result: %v\n", result)
	}
}
