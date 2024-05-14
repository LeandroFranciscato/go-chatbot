package rest

import (
	"review-chatbot/internal/usecase/chat"
	"review-chatbot/internal/usecase/customer"
	"review-chatbot/internal/usecase/flow"
	"review-chatbot/internal/usecase/order"

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
