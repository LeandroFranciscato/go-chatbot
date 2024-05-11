package router

import (
	"review-chatbot/internal/delivery/rest"
)

type router struct {
	rest.Server
}

func New(server rest.Server) router {
	return router{
		Server: server,
	}
}

func (router router) InitRoutes() {
	router.loadHTMLFiles()
	router.portaRoutes()
	router.chatRoutes()
}

func (router router) loadHTMLFiles() {
	files := []string{
		"files/home.html",
		"files/links.html",
		"files/form.html",
		"files/orders.html",
	}
	router.Engine.LoadHTMLFiles(files...)
}