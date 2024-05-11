package router

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (router router) portalRoutes() {
	router.publicRoutes()
	router.privateRoutes()
}

func (router router) publicRoutes() {
	router.loginRoutes()
	router.Engine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/home")
	})
}

func (router router) privateRoutes() {
	portalGroup := router.Engine.Group("/portal")

	portalGroup.Use(func(ctx *gin.Context) {
		customerID := sessions.Default(ctx).Get("customerID")
		if customerID == nil {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	})

	portalGroup.POST("/links", func(c *gin.Context) {
		c.HTML(http.StatusOK, "links.html", gin.H{})
	})

	router.OrderRoutes(portalGroup)
	router.reviewFlowRoute(portalGroup)
}
