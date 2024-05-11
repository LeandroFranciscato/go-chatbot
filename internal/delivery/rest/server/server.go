package server

import (
	"review-chatbot/internal/delivery/rest"
	"review-chatbot/internal/delivery/rest/router"
	"review-chatbot/internal/usecase/customer"
	"review-chatbot/internal/usecase/flow"
	"review-chatbot/internal/usecase/order"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Start(order order.Order, customer customer.Customer, flows ...flow.Flow) {
	store := cookie.NewStore([]byte("my-secret-key"))
	engine := gin.Default()
	engine.Use(sessions.Sessions("session", store))

	rest := rest.Server{
		Engine:   engine,
		Flows:    flows,
		Order:    order,
		Customer: customer,
	}

	router := router.New(rest)
	router.InitRoutes()

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
