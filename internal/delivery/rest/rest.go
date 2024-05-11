package rest

import (
	"review-chatbot/internal/usecase/flow"
	"review-chatbot/internal/usecase/order"

	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	Flows []flow.Flow
	Order order.Order
}
