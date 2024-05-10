package main

import (
	_ "embed"
	"review-chatbot/internal/delivery/rest"
	"review-chatbot/internal/usecase/flow"
)

//go:embed files/review.json
var reviewFlowJson []byte

func main() {

	usecase, err := flow.New(reviewFlowJson)
	if err != nil {
		panic(err)
	}

	rest.StartServer(usecase)
}
