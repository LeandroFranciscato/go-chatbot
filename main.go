package main

import (
	_ "embed"
	"review-chatbot/internal/repo"
)

//go:embed files/review.json
var reviewFlowJson []byte

func main() {

	repo.New()

	/*usecase, err := flow.New(reviewFlowJson)
	if err != nil {
		panic(err)
	}

	rest.StartServer(usecase)*/
}
