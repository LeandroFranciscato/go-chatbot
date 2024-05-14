package router

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (router router) routes() {
	router.publicRoutes()
	router.privateRoutes()
}

func (router router) publicRoutes() {
	router.login()
	router.home()
}

func (router router) privateRoutes() {
	portalGroup := router.Engine.Group("/portal")
	portalGroup.Use(router.authMiddleware)

	router.links(portalGroup)
	router.orderList(portalGroup)
	router.orderDelivered(portalGroup)

	chatGroup := portalGroup.Group("/chat")
	router.chat(chatGroup)
	router.chatReview(chatGroup)
	router.chatHistory(chatGroup)
	router.chatList(chatGroup)
}

func (router router) authMiddleware(ctx *gin.Context) {
	// validate if user is logged in
	customerID := sessions.Default(ctx).Get("customerID")
	if customerID == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	// validate if the logged user is the owner of the data
	customerIDParam := ctx.Param("customerID")
	if customerIDParam != "" && customerID != customerIDParam {
		ctx.AbortWithStatus(http.StatusForbidden)
	}
}
