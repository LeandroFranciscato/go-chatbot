package router

import (
	"html/template"
	"net/http"
	"review-chatbot/internal/usecase/flow"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (router router) chat(group *gin.RouterGroup) {
	group.POST("help", func(c *gin.Context) {
		router.chatHandler(c, router.ChatFlow)
	})
}

func (router router) chatHandler(c *gin.Context, flow flow.Flow) {
	// retrieve chat history
	customerID := sessions.Default(c).Get("customerID").(string)
	chatHistory, err := flow.GetHistory(c, customerID, "000000000000000000000000")
	if err != nil {
		c.String(http.StatusInternalServerError, "error finding chat history :"+err.Error())
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
	chatHistoryStr := c.Request.FormValue("history")
	if chatHistory.History != "" {
		chatHistoryStr = chatHistory.History
	}

	// identify if there is an user answer, avoid calling it in the initial step
	userAnswer := c.Request.FormValue("answer")
	timestamp := time.Now().Format(time.DateTime)
	nextStep := step
	botAnswer := ""
	if userAnswer != "" {
		chatHistoryStr += `<div class="user-message"><b>You:</b> ` + userAnswer + `<br><small>` + timestamp + `</small></div>`
		nextStep, botAnswer = flow.Answer(step, userAnswer)
		chatHistoryStr += `<div class="bot-message"><b>Bot:</b> ` + botAnswer + `<br><small>` + timestamp + `</small></div>`
	}

	// save the chat history
	err = flow.SaveHistory(c, nextStep, customerID, "000000000000000000000000", chatHistoryStr)
	if err != nil {
		c.String(http.StatusInternalServerError, "error registering chat :"+err.Error())
		return
	}

	//render chat form
	c.HTML(http.StatusOK, "chat.html", gin.H{
		"title":       flow.Name(),
		"step":        strconv.Itoa(nextStep),
		"customerID":  customerID,
		"final":       nextStep == flow.FinalStep(),
		"historyHTML": template.HTML(chatHistoryStr),
		"history":     chatHistoryStr,
	})
}
