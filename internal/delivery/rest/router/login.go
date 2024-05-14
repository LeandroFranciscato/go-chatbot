package router

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (router router) login() {

	router.Engine.POST("/login", func(ctx *gin.Context) {
		email := ctx.Request.FormValue("email")
		password := ctx.Request.FormValue("password")

		hash := md5.New()
		_, err := hash.Write([]byte(password))
		if err != nil {
			ctx.String(http.StatusInternalServerError, errors.New("error hashing password: "+err.Error()).Error())
			return
		}
		hashedPass := hex.EncodeToString(hash.Sum(nil))

		customer, err := router.Customer.Login(ctx, email, hashedPass)
		if err != nil {
			ctx.String(http.StatusInternalServerError, errors.New("error login in: "+err.Error()).Error())
			return
		}

		if customer.Email == "" {
			ctx.HTML(http.StatusUnauthorized, "home.html", gin.H{
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
