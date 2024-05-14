package router

import (
	"fmt"
	"net/http"
	"review-chatbot/internal/delivery/rest/dto"
	"review-chatbot/internal/domain/entity"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (router router) OrderRoutes(portalGroup *gin.RouterGroup) {

	portalGroup.GET("/orders", func(c *gin.Context) {

		customerID := sessions.Default(c).Get("customerID").(string)
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

		c.HTML(http.StatusOK, "orders.html", gin.H{
			"orders": ordersDto,
		})
	})

	portalGroup.POST("/customer/:customerID/order/:orderID/delivered", func(c *gin.Context) {

		customerID := c.Param("customerID")
		customerObjID, _ := primitive.ObjectIDFromHex(customerID)

		orderID := c.Param("orderID")
		orderObjID, _ := primitive.ObjectIDFromHex(orderID)

		order, err := router.Order.FindOne(c, customerObjID, orderObjID)
		if err != nil {
			c.String(http.StatusBadRequest, "error finding order: "+err.Error())
			return
		}

		order.Status = entity.OrderStatusDelivered
		if err = router.Order.UpdateOne(c, order); err != nil {
			c.String(http.StatusInternalServerError, "error updating order: "+err.Error())
			return
		}

		reviewFlowRoute := fmt.Sprintf("/portal/chat/review/customer/%s/order/%s", customerID, orderID)
		c.Redirect(http.StatusPermanentRedirect, reviewFlowRoute)
	})

}
