package rest

import (
	"github.com/LeandroFranciscato/go-chatbot/internal/usecase/chat"
	"github.com/LeandroFranciscato/go-chatbot/internal/usecase/customer"
	"github.com/LeandroFranciscato/go-chatbot/internal/usecase/flow"
	"github.com/LeandroFranciscato/go-chatbot/internal/usecase/order"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	ReviewFlow flow.Flow
	ChatFlow   flow.Flow
	Order      order.Order
	Customer   customer.Customer
	Chat       chat.Chat
}
