package controllers

import (
	"log"
	"net/http"
	"session/database"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// todo logger middleware
func Logger() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("router:", ctx.Request.URL)
		ctx.Next()
	}
}

var store = sessions.NewCookieStore([]byte("super-secret"))

// todo middleware to check current session
func Session() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session, _ := store.Get(ctx.Request, "session-name")
		username, ok := session.Values["username"].(string)
		if !ok {
			log.Println("need session!")
			ctx.Redirect(http.StatusSeeOther, "/login")
			return
		}

		// send data to the next route
		ctx.Set("username", username)
		ctx.Next()
	}
}

// todo templating for login
func Login(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", nil)
}

// todo redirect to login html
func Main(ctx *gin.Context) {
	// get the data from middleware
	username := ctx.MustGet("username").(string)
	ctx.HTML(http.StatusOK, "index.html", gin.H{"Message": username})
}

// todo create a session
func LoginPostHandler(ctx *gin.Context) {
	var user database.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	database.DB.First(&user, &user.Username)

	session, _ := store.Get(ctx.Request, "session-name")
	session.Values["username"] = &user.Username
	session.Save(ctx.Request, ctx.Writer)

	ctx.Redirect(http.StatusSeeOther, "/")
}
