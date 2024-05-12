package router

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (router router) reviewFlowRoute(portalGroup *gin.RouterGroup) {

	chatGroup := portalGroup.Group("/chat")

	chatGroup.POST("review/customer/:customerID/order/:orderID", func(c *gin.Context) {
		customerID := c.Param("customerID")
		orderID := c.Param("orderID")

		var err error
		step := 1
		stepStr := c.Request.FormValue("step")
		if stepStr != "" {
			step, err = strconv.Atoi(stepStr)
			if err != nil {
				c.String(http.StatusBadRequest, "error parsing step :"+err.Error())
				return
			}
		}

		nextStep := step
		answer := ""
		userAnswer := c.Request.FormValue("answer")
		if userAnswer != "" {
			nextStep, answer = router.ReviewFlow.Answer(step, userAnswer)
		}

		question, err := router.ReviewFlow.Ask(customerID, orderID, nextStep)
		if err != nil {
			c.String(http.StatusInternalServerError, "error asking :"+err.Error())
			return
		}

		c.HTML(http.StatusOK, "chat.html", gin.H{
			"title":      router.ReviewFlow.Name(),
			"answer":     answer,
			"question":   question,
			"step":       strconv.Itoa(nextStep),
			"customerID": customerID,
			"orderID":    orderID,
		})
	})

}
