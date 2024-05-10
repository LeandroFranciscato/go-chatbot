package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (rest rest) InitRoutes() {
	rest.LoadHTMLFiles()
	rest.portaRoutes()
	rest.chatRoutes()
}

func (rest rest) LoadHTMLFiles() {
	rest.Engine.LoadHTMLFiles("files/home.html", "files/links.html", "files/form.html")
}

func (rest rest) portaRoutes() {
	portalGroup := rest.Engine.Group("/portal")
	portalGroup.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	portalGroup.POST("/links", func(c *gin.Context) {
		c.HTML(http.StatusOK, "links.html", gin.H{})
	})

	rest.Engine.GET("/", func(ctx *gin.Context) {
		ctx.Redirect(http.StatusPermanentRedirect, "/portal/home")
	})
}

func (rest rest) chatRoutes() {

	chatGroup := rest.Engine.Group("/chat")

	for _, flow := range rest.flows {

		chatGroup.GET("/"+flow.ID(), func(c *gin.Context) {
			step := 1
			c.HTML(http.StatusOK, "form.html", gin.H{
				"title":    "Review",
				"question": flow.Ask(step),
				"step":     strconv.Itoa(step),
			})
		})

		chatGroup.POST("/"+flow.ID(), func(c *gin.Context) {
			step, _ := strconv.Atoi(c.Request.FormValue("step"))
			userAnswer := c.Request.FormValue("answer")
			nextStep, answer := flow.Answer(step, userAnswer)
			c.HTML(http.StatusOK, "form.html", gin.H{
				"title":    "Review",
				"answer":   answer,
				"question": flow.Ask(nextStep),
				"step":     strconv.Itoa(nextStep),
			})
		})
	}
}
