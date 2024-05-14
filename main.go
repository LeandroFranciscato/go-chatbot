package main

import (
	"context"
	_ "embed"
	"os"
	"review-chatbot/internal/delivery/rest/server"
	"review-chatbot/internal/domain/entity"
	"review-chatbot/internal/repo"
	"review-chatbot/internal/usecase/chat"
	"review-chatbot/internal/usecase/customer"
	"review-chatbot/internal/usecase/flow"
	"review-chatbot/internal/usecase/order"
	"review-chatbot/internal/util"
)

//go:embed files/review.json
var reviewFlowJson []byte

//go:embed files/chat.json
var chatJson []byte

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

	customerRepo, err := repo.New[entity.Customer](mongoUri, mongoUser, mongoPass, mongoDb, entity.Customer{}.GetCollectionName())
	if err != nil {
		panic(err)
	}
	customerUsercase := customer.New(customerRepo)

	orderRepo, err := repo.New[entity.Order](mongoUri, mongoUser, mongoPass, mongoDb, entity.Order{}.GetCollectionName())
	if err != nil {
		panic(err)
	}

	orderUsecase := order.New(orderRepo)

	chatRepo, err := repo.New[entity.Chat](mongoUri, mongoUser, mongoPass, mongoDb, entity.Chat{}.GetCollectionName())
	if err != nil {
		panic(err)
	}

	chatUsecase := chat.New(chatRepo)

	reviewFlowUsecase, err := flow.New(reviewFlowJson, chatRepo)
	if err != nil {
		panic(err)
	}

	chatFlowUsecase, err := flow.New(chatJson, chatRepo)
	if err != nil {
		panic(err)
	}

	server.Start(orderUsecase, customerUsercase, chatUsecase, reviewFlowUsecase, chatFlowUsecase)
}
