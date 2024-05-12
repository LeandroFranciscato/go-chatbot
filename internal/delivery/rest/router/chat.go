package router

import (
	"encoding/json"
	"html/template"
	"net/http"
	"review-chatbot/internal/domain/entity"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (router router) reviewFlowRoute(portalGroup *gin.RouterGroup) {

	chatGroup := portalGroup.Group("/chat")

	chatGroup.POST("review/customer/:customerID/order/:orderID", func(c *gin.Context) {
		// identify the step user is
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

		// logic to retrieve order from db only in the first step, then reuse it
		orderStr := c.Request.FormValue("order")
		customerID := c.Param("customerID")
		orderID := c.Param("orderID")
		var order entity.Order
		if orderStr == "" {
			customerObjID, _ := primitive.ObjectIDFromHex(customerID)
			orderObjID, _ := primitive.ObjectIDFromHex(orderID)

			order, err = router.Order.FindOne(c, customerObjID, orderObjID)
			if err != nil {
				c.String(http.StatusInternalServerError, "error finding order :"+err.Error())
				return
			}
			orderBytes, _ := json.Marshal(order)
			orderStr = string(orderBytes)
		} else {
			err = json.Unmarshal([]byte(orderStr), &order)
			if err != nil {
				c.String(http.StatusInternalServerError, "error unmarshalling order :"+err.Error())
				return
			}
		}

		// identify if there is an user answer (avoid calling in first step)
		history := c.Request.FormValue("history")
		nextStep := step
		botAnswer := ""
		userAnswer := c.Request.FormValue("answer")
		if userAnswer != "" {
			history += `<div class="user-message"><b>You:</b> ` + userAnswer + `</div>`
			nextStep, botAnswer = router.ReviewFlow.Answer(step, userAnswer)
			history += `<div class="bot-message"><b>Bot:</b> ` + botAnswer + `</div>`
		}

		// ask the user the next botQuestion
		botQuestion, err := router.ReviewFlow.Ask(order, nextStep)
		if err != nil {
			c.String(http.StatusInternalServerError, "error asking :"+err.Error())
			return
		}
		history += `<div class="bot-message"><b>Bot:</b> ` + string(botQuestion) + `</div>`

		//render chat form
		c.HTML(http.StatusOK, "chat.html", gin.H{
			"title":        router.ReviewFlow.Name(),
			"bot_answer":   botAnswer,
			"bot_question": botQuestion,
			"step":         strconv.Itoa(nextStep),
			"customerID":   customerID,
			"orderID":      orderID,
			"final":        nextStep == router.ReviewFlow.FinalStep(),
			"order":        orderStr,
			"historyHTML":  template.HTML(history),
			"history":      history,
		})
	})

}
