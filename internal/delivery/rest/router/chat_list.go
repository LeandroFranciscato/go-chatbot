package router

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (router router) chatList(group *gin.RouterGroup) {

	group.GET("list", func(c *gin.Context) {
		customerID := sessions.Default(c).Get("customerID").(string)
		router.chatListHandler(c, customerID)
	})
}

func (router router) chatListHandler(c *gin.Context, customerID string) {
	customerObjID, _ := primitive.ObjectIDFromHex(customerID)

	chatList, err := router.Chat.FindByCustomer(c, customerObjID)
	if err != nil {
		c.String(http.StatusBadRequest, "error finding chat list: "+err.Error())
		return
	}

	c.HTML(http.StatusOK, "chat_list.html", gin.H{
		"chatList": chatList,
	})
}
