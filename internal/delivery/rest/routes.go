package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (rest rest) InitRoutes() {
	rest.LoadHTMLFiles()
	rest.portaRoutes()
	rest.chatRoutes()
}

func (rest rest) LoadHTMLFiles() {
	rest.Engine.LoadHTMLFiles("files/home.html", "files/links.html", "files/form.html", "files/orders.html")
}

func (rest rest) portaRoutes() {
	portalGroup := rest.Engine.Group("/portal")
	portalGroup.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	portalGroup.POST("/links", func(c *gin.Context) {
		c.HTML(http.StatusOK, "links.html", gin.H{})
	})

	portalGroup.GET("/orders", func(c *gin.Context) {
		customer, _ := primitive.ObjectIDFromHex("663eb8bff247a62df85b0ae1") // must change to the logged customer
		orders, _ := rest.order.FindByCustomer(c, customer)
		c.HTML(http.StatusOK, "orders.html", gin.H{
			"orders": orders,
		})
	})

	rest.Engine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/portal/home")
	})
}

func (rest rest) chatRoutes() {

	chatGroup := rest.Engine.Group("/chat")

	for _, flow := range rest.flows {

		chatGroup.GET("/"+flow.ID(), func(c *gin.Context) {
			step := 1
			c.HTML(http.StatusOK, "form.html", gin.H{
				"title":    "Review",
				"question": flow.Ask(step),
				"step":     strconv.Itoa(step),
			})
		})

		chatGroup.POST("/"+flow.ID(), func(c *gin.Context) {
			step, _ := strconv.Atoi(c.Request.FormValue("step"))
			userAnswer := c.Request.FormValue("answer")
			nextStep, answer := flow.Answer(step, userAnswer)
			c.HTML(http.StatusOK, "form.html", gin.H{
				"title":    "Review",
				"answer":   answer,
				"question": flow.Ask(nextStep),
				"step":     strconv.Itoa(nextStep),
			})
		})
	}
}
