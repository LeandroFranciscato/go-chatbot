package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router router) home() {

	router.Engine.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	router.Engine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/home")
	})

}
