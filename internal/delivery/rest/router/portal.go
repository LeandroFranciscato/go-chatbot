package router

import (
	"net/http"

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
		customer, _ := primitive.ObjectIDFromHex("663eb8bff247a62df85b0ae1") // must change to the logged customer
		orders, _ := router.Order.FindByCustomer(c, customer)
		c.HTML(http.StatusOK, "orders.html", gin.H{
			"orders": orders,
		})
	})

	router.Engine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/portal/home")
	})
}
