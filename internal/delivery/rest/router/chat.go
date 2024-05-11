package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (router router) chatRoutes() {

	chatGroup := router.Engine.Group("/chat")

	for _, flow := range router.Flows {

		chatGroup.POST("/"+flow.ID(), func(c *gin.Context) {
			var err error
			step := 1
			stepStr := c.Request.FormValue("step")
			if stepStr != "" {
				step, err = strconv.Atoi(stepStr)
				if err != nil {
					c.String(http.StatusBadRequest, "error parsing step :"+err.Error())
				}
			}

			nextStep := step
			answer := ""
			userAnswer := c.Request.FormValue("answer")
			if userAnswer != "" {
				nextStep, answer = flow.Answer(step, userAnswer)
			}
			c.HTML(http.StatusOK, "form.html", gin.H{
				"title":    flow.Name(),
				"answer":   answer,
				"question": flow.Ask(nextStep),
				"step":     strconv.Itoa(nextStep),
			})
		})
	}
}
