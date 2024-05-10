package main

import (
	"context"
	_ "embed"
	"fmt"
	"review-chatbot/internal/domain/entity"
	"review-chatbot/internal/repo"
)

//go:embed files/review.json
var reviewFlowJson []byte

func main() {

	repo, err := repo.New[entity.Customer]("mongodb://localhost:27017", "root", "example", "chatbot", "customer")
	if err != nil {
		panic(err)
	}

	res, err := repo.List(context.Background())
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
