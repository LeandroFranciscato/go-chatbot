package router

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (router router) home() {

	router.Engine.GET("/home", func(c *gin.Context) {
		customerID := sessions.Default(c).Get("customerID")
		if customerID != nil {
			c.HTML(http.StatusOK, "links.html", gin.H{})
			return
		}
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	router.Engine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/home")
	})

}
