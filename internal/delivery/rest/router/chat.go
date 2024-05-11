package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (router router) chatRoutes() {

	chatGroup := router.Engine.Group("/chat")

	for _, flow := range router.Flows {

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
