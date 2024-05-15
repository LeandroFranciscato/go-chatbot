package server

import (
	"github.com/LeandroFranciscato/go-chatbot/internal/delivery/rest"
	"github.com/LeandroFranciscato/go-chatbot/internal/delivery/rest/router"
	"github.com/LeandroFranciscato/go-chatbot/internal/usecase/chat"
	"github.com/LeandroFranciscato/go-chatbot/internal/usecase/customer"
	"github.com/LeandroFranciscato/go-chatbot/internal/usecase/flow"
	"github.com/LeandroFranciscato/go-chatbot/internal/usecase/order"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Start(order order.Order, customer customer.Customer, Chat chat.Chat, reviewFlow flow.Flow, chatFlow flow.Flow) {
	store := cookie.NewStore([]byte("my-secret-key"))
	engine := gin.Default()
	engine.Use(sessions.Sessions("session", store))

	rest := rest.Server{
		Engine:     engine,
		ReviewFlow: reviewFlow,
		Order:      order,
		Customer:   customer,
		Chat:       Chat,
		ChatFlow:   chatFlow,
	}

	router := router.New(rest)
	router.InitRoutes()

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
