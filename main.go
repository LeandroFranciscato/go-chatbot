package main

import (
	"context"
	_ "embed"
	"fmt"
	"review-chatbot/internal/domain/entity"
	"review-chatbot/internal/repo"

	"go.mongodb.org/mongo-driver/bson"
)

//go:embed files/review.json
var reviewFlowJson []byte

func main() {

	repo, err := repo.New[entity.Customer]("mongodb://localhost:27017", "root", "example", "chatbot", "customer")
	if err != nil {
		panic(err)
	}

	res, err := repo.FindOne(context.Background(),
		bson.D{
			{Key: "$and", Value: bson.A{
				bson.D{{Key: "email", Value: bson.D{{Key: "$eq", Value: "c1@gmail.com"}}}},
				bson.D{{Key: "password", Value: bson.D{{Key: "$eq", Value: "pass"}}}},
			}},
		},
	)
	if err != nil {
		panic(err)
	}

	fmt.Printf("res: %v\n", res)

	/*usecase, err := flow.New(reviewFlowJson)
	if err != nil {
		panic(err)
	}

	rest.StartServer(usecase)*/
}
