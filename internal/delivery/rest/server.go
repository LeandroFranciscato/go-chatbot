package rest

import (
	"review-chatbot/internal/usecase/flow"
	"review-chatbot/internal/usecase/order"

	"github.com/gin-gonic/gin"
)

type rest struct {
	*gin.Engine
	flows []flow.Flow
	order order.Order
}

func StartServer(order order.Order, flows ...flow.Flow) {
	engine := gin.Default()

	rest := rest{
		Engine: engine,
		flows:  flows,
		order:  order,
	}

	rest.InitRoutes()

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
