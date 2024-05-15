package router

import (
	"net/http"
	"review-chatbot/internal/delivery/rest/dto"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (router router) orderList(group *gin.RouterGroup) {

	group.GET("order/list", func(c *gin.Context) {
		customerID := sessions.Default(c).Get("customerID").(string)
		router.orderListHandler(c, customerID)
	})
}

func (router router) orderListHandler(c *gin.Context, customerID string) {
	customerObjID, _ := primitive.ObjectIDFromHex(customerID)
	orders, err := router.Order.FindByCustomer(c, customerObjID)
	if err != nil {
		c.String(http.StatusBadRequest, "error finding orders: "+err.Error())
		return
	}

	var ordersDto []dto.Order
	for _, order := range orders {
		orderDto := dto.Order{Order: order}
		chatHistory, err := router.ReviewFlow.GetHistory(c, customerID, order.ID.Hex())
		if err != nil {
			c.String(http.StatusBadRequest, "error finding chat history: "+err.Error())
			return
		}
		orderDto.ChatStatus = chatHistory.Status
		ordersDto = append(ordersDto, orderDto)
	}

	c.HTML(http.StatusOK, "order_list.html", gin.H{
		"orders": ordersDto,
	})
}
