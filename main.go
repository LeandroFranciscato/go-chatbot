package main

import (
	"context"
	_ "embed"
	"os"
	"review-chatbot/internal/delivery/rest/server"
	"review-chatbot/internal/domain/entity"
	"review-chatbot/internal/repo"
	"review-chatbot/internal/usecase/flow"
	"review-chatbot/internal/usecase/order"
	"review-chatbot/internal/util"
)

//go:embed files/review.json
var reviewFlowJson []byte

func main() {

	mongoUri := "mongodb://localhost:27017"
	mongoUser := "root"
	mongoPass := "example"
	mongoDb := "chatbot"

	migrate := os.Getenv("MIGRATE")
	if migrate != "" {
		err := util.Migrate(context.Background(), mongoUri, mongoUser, mongoPass, mongoDb)
		if err != nil {
			panic(err)
		}
	}
	/*
		repo, err := repo.New[entity.Order]("mongodb://localhost:27017", "root", "example", "chatbot", "order")
		if err != nil {
			panic(err)
		}

		id, err := primitive.ObjectIDFromHex("663ed9ec2c2b9af97e15ccc8")
		if err != nil {
			panic(err)
		}

		one, err := repo.FindOne(context.Background(), bson.D{{Key: "_id", Value: id}})
		if err != nil {
			panic(err)
		}
		one.Status = entity.OrderStatusDelivered

		usecase := order.New(repo)

		err = usecase.Update(context.Background(), one)
		if err != nil {
			panic(err)
		}
	*/
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

		res, err := repo.Find(context.Background(),
			bson.D{},
		)
		if err != nil {
			panic(err)
		}

		fmt.Printf("res: %v\n", res)
	*/

	orderRepo, err := repo.New[entity.Order]("mongodb://localhost:27017", "root", "example", "chatbot", "order")
	if err != nil {
		panic(err)
	}
	orderUsecase := order.New(orderRepo)

	usecase, err := flow.New(reviewFlowJson)
	if err != nil {
		panic(err)
	}

	server.Start(orderUsecase, usecase)
}
