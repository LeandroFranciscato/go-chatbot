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
	})

	portalGroup.POST("/links", func(c *gin.Context) {
		c.HTML(http.StatusOK, "links.html", gin.H{})
	})

	router.OrderRoutes(portalGroup)
	router.reviewFlowRoute(portalGroup)
}
