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
	/*
			resID, err := repo.InsertOne(context.Background(), entity.Customer{Name: "Giovana", Email: "gihromeu@gmail.com", Password: "pass"})
			if err != nil {
				panic(err)
			}

			fmt.Printf("resID: %v\n", resID)

		res, err := repo.FindOne(context.Background(),
			bson.D{
				{Key: "$and", Value: bson.A{
					bson.D{{Key: "email", Value: bson.D{{Key: "$eq", Value: "lbfranciscato@gmail.com"}}}},
					bson.D{{Key: "password", Value: bson.D{{Key: "$eq", Value: "pass"}}}},
				}},
			},
		)
		if err != nil {
			panic(err)
		}
	*/
	res, err := repo.Find(context.Background(),
		bson.D{},
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
