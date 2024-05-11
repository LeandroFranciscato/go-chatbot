package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (router router) reviewFlowRoute(portalGroup *gin.RouterGroup) {

	chatGroup := portalGroup.Group("/chat")

	chatGroup.POST("review/customer/:customerID/order/:orderID", func(c *gin.Context) {
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
			nextStep, answer = router.ReviewFlow.Answer(step, userAnswer)
		}
		c.HTML(http.StatusOK, "chat.html", gin.H{
			"title":      router.ReviewFlow.Name(),
			"answer":     answer,
			"question":   router.ReviewFlow.Ask(nextStep),
			"step":       strconv.Itoa(nextStep),
			"customerID": c.Param("customerID"),
			"orderID":    c.Param("orderID"),
		})
	})

}
