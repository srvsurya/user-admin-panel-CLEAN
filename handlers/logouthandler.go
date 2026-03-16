package handlers

import ("github.com/gin-contrib/sessions"
"github.com/gin-gonic/gin")

func LogoutHandler()gin.HandlerFunc{return func(c *gin.Context){
	session:=sessions.Default(c)
	session.Clear()
	session.Save()

	c.Redirect(302,"/login")
}}