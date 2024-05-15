package router

import (
	"fmt"
	"net/http"

	"github.com/LeandroFranciscato/go-chatbot/internal/domain/entity"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (router router) orderDelivered(group *gin.RouterGroup) {

	group.POST("/customer/:customerID/order/:orderID/delivered", func(c *gin.Context) {

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
