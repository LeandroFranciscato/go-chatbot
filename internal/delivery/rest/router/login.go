package router

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (router router) loginRoutes() {

	router.Engine.GET("/home", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{})
	})

	router.Engine.POST("/login", func(ctx *gin.Context) {
		email := ctx.Request.FormValue("email")
		password := ctx.Request.FormValue("password")

		customer, err := router.Customer.Login(ctx, email, password)
		if err != nil {
			ctx.String(http.StatusInternalServerError, errors.New("error login in: "+err.Error()).Error())
			return
		}

		if customer.Email == "" {
			ctx.HTML(http.StatusOK, "home.html", gin.H{
				"message": "email/password incorrect",
			})
			return
		}

		session := sessions.Default(ctx)
		session.Set("customerID", customer.ID.Hex())
		session.Save()

		ctx.Redirect(http.StatusPermanentRedirect, "/portal/links")
	})

}
