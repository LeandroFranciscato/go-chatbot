package router

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router router) chatHistory(group *gin.RouterGroup) {

	group.POST("customer/:customerID/order/:orderID", func(c *gin.Context) {
		customerID := c.Param("customerID")
		orderID := c.Param("orderID")

		// retrieve chat chatHistory
		chatHistory, err := router.ReviewFlow.GetHistory(c, customerID, orderID)
		if err != nil {
			c.String(http.StatusInternalServerError, "error finding chat history :"+err.Error())
			return
		}

		//render chat form
		c.HTML(http.StatusOK, "chat.html", gin.H{
			"title":       router.ReviewFlow.Name(),
			"historyHTML": template.HTML(chatHistory.History),
			"readonly":    true,
		})
	})

}
