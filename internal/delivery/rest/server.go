package rest

import (
	"review-chatbot/internal/usecase/flow"

	"github.com/gin-gonic/gin"
)

type rest struct {
	*gin.Engine
	flows []flow.Flow
}

func StartServer(flows ...flow.Flow) {
	engine := gin.Default()

	rest := rest{
		Engine: engine,
		flows:  flows,
	}

	rest.InitRoutes()

	if err := engine.Run(); err != nil {
		panic(err)
	}
}
