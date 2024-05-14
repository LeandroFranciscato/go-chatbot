package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router router) links(group *gin.RouterGroup) {
	group.POST("/links", func(c *gin.Context) {
		c.HTML(http.StatusOK, "links.html", gin.H{})
	})
}
