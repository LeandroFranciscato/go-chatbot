package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
	"review-chatbot/internal/domain/entity"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (router router) chat(group *gin.RouterGroup) {

	group.POST("review/customer/:customerID/order/:orderID", func(c *gin.Context) {
		// retrieve chat history
		customerID := c.Param("customerID")
		orderID := c.Param("orderID")
		chatHistory, err := router.ReviewFlow.GetHistory(c, customerID, orderID)
		if err != nil {
			c.String(http.StatusInternalServerError, "error finding chat history :"+err.Error())
			return
		}

		// retrieve order
		orderStr, order, err := router.getOrder(c, customerID, orderID)
		if err != nil {
			c.String(http.StatusInternalServerError, "error finding order :"+err.Error())
			return
		}

		// identify the step user is
		step := 1
		stepStr := c.Request.FormValue("step")
		if stepStr != "" {
			step, err = strconv.Atoi(stepStr)
			if err != nil {
				c.String(http.StatusBadRequest, "error parsing step :"+err.Error())
				return
			}
		}

		// if user is in a step ahead, just render the chat form
		if chatHistory.CurrentStep > step {
			c.HTML(http.StatusOK, "chat.html", gin.H{
				"title":       router.ReviewFlow.Name(),
				"step":        strconv.Itoa(chatHistory.CurrentStep),
				"customerID":  customerID,
				"orderID":     orderID,
				"final":       chatHistory.CurrentStep == router.ReviewFlow.FinalStep(),
				"order":       orderStr,
				"historyHTML": template.HTML(chatHistory.History),
				"history":     chatHistory.History,
			})
			return
		}

		// identify if there is an user answer, avoid calling it in the initial step
		chatHistoryStr := c.Request.FormValue("history")
		userAnswer := c.Request.FormValue("answer")
		timestamp := time.Now().Format(time.DateTime)
		nextStep := step
		botAnswer := ""
		if userAnswer != "" {
			chatHistoryStr += `<div class="user-message"><b>You:</b> ` + userAnswer + `<br><small>` + timestamp + `</small></div>`
			nextStep, botAnswer = router.ReviewFlow.Answer(step, userAnswer)
			chatHistoryStr += `<div class="bot-message"><b>Bot:</b> ` + botAnswer + `<br><small>` + timestamp + `</small></div>`
		}

		// ask the user the next botQuestion
		botQuestion := router.ReviewFlow.Ask(nextStep)
		if err != nil {
			c.String(http.StatusInternalServerError, "error asking :"+err.Error())
			return
		}

		// replace order variables
		botQuestionReplaced, err := router.replaceOrderVariables(botQuestion, order)
		if err != nil {
			c.String(http.StatusInternalServerError, "error replacing order variables :"+err.Error())
			return
		}
		chatHistoryStr += `<div class="bot-message"><b>Bot:</b> ` + botQuestionReplaced + `<br><small>` + timestamp + `</small></div>`

		// save the chat history
		err = router.ReviewFlow.SaveHistory(c, nextStep, customerID, orderID, chatHistoryStr)
		if err != nil {
			c.String(http.StatusInternalServerError, "error registering chat :"+err.Error())
			return
		}

		//render chat form
		c.HTML(http.StatusOK, "chat.html", gin.H{
			"title":       router.ReviewFlow.Name(),
			"step":        strconv.Itoa(nextStep),
			"customerID":  customerID,
			"orderID":     orderID,
			"final":       nextStep == router.ReviewFlow.FinalStep(),
			"order":       orderStr,
			"historyHTML": template.HTML(chatHistoryStr),
			"history":     chatHistoryStr,
		})
	})
}

// getOrder retrieves order from db only in the first step, then reuse it
func (router router) getOrder(c *gin.Context, customerID string, orderID string) (orderStr string, order entity.Order, err error) {

	if orderStr == "" {
		customerObjID, _ := primitive.ObjectIDFromHex(customerID)
		orderObjID, _ := primitive.ObjectIDFromHex(orderID)

		order, err = router.Order.FindOne(c, customerObjID, orderObjID)
		if err != nil {
			return "", entity.Order{}, errors.New("error finding order: " + err.Error())
		}

		orderBytes, _ := json.Marshal(order)
		orderStr = string(orderBytes)
	} else {
		err = json.Unmarshal([]byte(orderStr), &order)
		if err != nil {
			return "", entity.Order{}, errors.New("error unmarshalling order: " + err.Error())
		}
	}

	return orderStr, order, nil
}

// replaceOrderVariables replaces order variables in the botQuestion html
func (router router) replaceOrderVariables(botQuestion string, order entity.Order) (string, error) {
	tmpl, err := template.New("question").Parse(botQuestion)
	if err != nil {
		return "", errors.New("error creating templace :" + err.Error())
	}

	var botQuestionBuff bytes.Buffer
	err = tmpl.Execute(&botQuestionBuff, order)
	if err != nil {
		return "", errors.New("error executing template :" + err.Error())
	}
	return botQuestionBuff.String(), nil
}
