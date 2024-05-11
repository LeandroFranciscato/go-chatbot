package router

import (
	"fmt"
	"net/http"
	"review-chatbot/internal/domain/entity"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (router router) OrderRoutes(portalGroup *gin.RouterGroup) {

	portalGroup.GET("/orders", func(c *gin.Context) {

		customerStrID := sessions.Default(c).Get("customerID").(string)
		customerID, err := primitive.ObjectIDFromHex(customerStrID)
		if err != nil {
			c.String(http.StatusBadRequest, "error parsing customer obj id: "+err.Error())
			return
		}

		orders, _ := router.Order.FindByCustomer(c, customerID)
		c.HTML(http.StatusOK, "orders.html", gin.H{
			"orders": orders,
		})
	})

	portalGroup.POST("/customer/:customerID/order/:orderID/delivered", func(c *gin.Context) {

		customerID := c.Param("customerID")
		sessionCustomerID := sessions.Default(c).Get("customerID").(string)
		if customerID != sessionCustomerID {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		customerObjID, err := primitive.ObjectIDFromHex(customerID)
		if err != nil {
			c.String(http.StatusBadRequest, "error parsing customer object id: "+err.Error())
			return
		}

		orderID := c.Param("orderID")
		orderObjID, err := primitive.ObjectIDFromHex(orderID)
		if err != nil {
			c.String(http.StatusBadRequest, "error parsing order object id: "+err.Error())
			return
		}

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
