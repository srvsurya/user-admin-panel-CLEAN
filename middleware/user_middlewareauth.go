package middleware

import ("github.com/gin-contrib/sessions"
"github.com/gin-gonic/gin")

func UserMiddlewareAuth()gin.HandlerFunc{return func(c *gin.Context){
	session:=sessions.Default(c)
	user_id:=session.Get("user_id")

	if user_id ==nil{
		c.Redirect(302,"/login")
		c.Abort()
		return
	}
	c.Next()

}}