package middleware

import ("github.com/gin-contrib/sessions"
"github.com/gin-gonic/gin")

func AdminMiddlewareAuth()gin.HandlerFunc{return func(c *gin.Context){
	session:=sessions.Default(c)
	role:=session.Get("role")

	if role!="admin"{
		c.AbortWithStatus(403)
		return
	}
	c.Next()
}}