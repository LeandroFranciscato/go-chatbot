package router

import (
	"net/http"
	"review-chatbot/internal/domain/entity"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (router router) portaRoutes() {
	portalGroup := router.Engine.Group("/portal")
	portalGroup.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	portalGroup.POST("/links", func(c *gin.Context) {
		c.HTML(http.StatusOK, "links.html", gin.H{})
	})

	portalGroup.GET("/orders", func(c *gin.Context) {
		customer, _ := primitive.ObjectIDFromHex("663ed9ec2c2b9af97e15ccc1") // must change to the logged customer
		orders, _ := router.Order.FindByCustomer(c, customer)
		c.HTML(http.StatusOK, "orders.html", gin.H{
			"orders": orders,
		})
	})

	portalGroup.POST("/order/:id/delivered", func(c *gin.Context) {
		orderID := c.Param("id")
		objectId, err := primitive.ObjectIDFromHex(orderID)
		if err != nil {
			c.String(http.StatusBadRequest, "error parsing object id: "+err.Error())
			return
		}
		order, err := router.Order.FindOne(c, objectId)
		if err != nil {
			c.String(http.StatusBadRequest, "error finding order: "+err.Error())
			return
		}

		order.Status = entity.OrderStatusDelivered
		if err = router.Order.UpdateOne(c, order); err != nil {
			c.String(http.StatusInternalServerError, "error updating order: "+err.Error())
		}

		c.Redirect(http.StatusPermanentRedirect, "/chat/review")
	})

	router.Engine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/portal/home")
	})
}
