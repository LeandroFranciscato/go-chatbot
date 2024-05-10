package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (rest rest) InitRoutes() {

	api := rest.Engine.Group("/api")
	api.GET("customer/:customer/order/:order/delivery-confirmed", func(c *gin.Context) {
		c.Redirect(http.StatusPermanentRedirect, "/chat/review/")
	})

	rest.Engine.LoadHTMLFiles("files/form.html")

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
