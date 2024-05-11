package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router router) portalRoutes() {
	portalGroup := router.Engine.Group("/portal")
	portalGroup.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	portalGroup.POST("/links", func(c *gin.Context) {
		c.HTML(http.StatusOK, "links.html", gin.H{})
	})

	router.Engine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/portal/home")
	})

	router.loginRoutes(portalGroup)
	router.OrderRoutes(portalGroup)
}
