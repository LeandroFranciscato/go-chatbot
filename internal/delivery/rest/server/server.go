package server

import (
	"review-chatbot/internal/delivery/rest"
	"review-chatbot/internal/delivery/rest/router"
	"review-chatbot/internal/usecase/flow"
	"review-chatbot/internal/usecase/order"

	"github.com/gin-gonic/gin"
)

func Start(order order.Order, flows ...flow.Flow) {
	engine := gin.Default()

	rest := rest.Server{
		Engine: engine,
		Flows:  flows,
		Order:  order,
	}

	router := router.New(rest)
	router.InitRoutes()

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
