package router

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (router router) logout(group *gin.RouterGroup) {

	group.GET("logout", func(ctx *gin.Context) {

		session := sessions.Default(ctx)
		session.Clear()
		err := session.Save()
		if err != nil {
			ctx.String(http.StatusInternalServerError, errors.New("error saving session: "+err.Error()).Error())
			return
		}

		customerID := sessions.Default(ctx).Get("customerID")
		fmt.Printf("customerID: %v\n", customerID)

		ctx.HTML(http.StatusOK, "home.html", gin.H{})
	})

}
