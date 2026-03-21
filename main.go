package main

import (
	"Week_12/db"
	"Week_12/handlers"
	"Week_12/middleware"
	"net/http"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"Week_12/logger"

)

func main(){
	logger.Init()
	defer logger.Log.Sync()
	database:=db.Connect()
	db.Migrate(database)
	r:=gin.Default()
	//template functions
	r.SetFuncMap(template.FuncMap{
		"add":func( a int, b int )int{return a+b},
		"sub":func(a int, b int)int{return a-b},
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/static","./static")

	store:=cookie.NewStore([]byte("super-secret-key"))
		store.Options(sessions.Options{
		Path:     "/",
		MaxAge:   3600, 
		HttpOnly: true,
		Secure:   false, 
		SameSite: http.SameSiteLaxMode,
	})
	r.Use(sessions.Sessions("mysession",store))
	r.Use(middleware.RequestIDCreator())

	r.GET("/",func(c *gin.Context){
		session:=sessions.Default(c)
		if session.Get("user_id")!=nil{
				if session.Get("role") == "admin"{
					c.Redirect(302,"/adminpanel")
					return
				}else{
					c.Redirect(302,"/home")
					return
				}
				
		}

		c.HTML(200,"login.html",nil)
	})
	r.GET("/signup",func(c *gin.Context){
		c.HTML(200,"sign_up.html",nil)
	})
	r.GET("/home",middleware.UserMiddlewareAuth(),handlers.HomeHandler(database))
	r.GET("/login",func(c *gin.Context){
		c.HTML(200,"login.html",nil)
	})
	r.GET("/adminpanel",middleware.UserMiddlewareAuth(),middleware.AdminMiddlewareAuth(),handlers.AdminviewHandler(database))
	r.GET("/admin/create",middleware.UserMiddlewareAuth(),middleware.AdminMiddlewareAuth(),func(c *gin.Context){
		c.HTML(200,"create_user.html",nil)
	})
	r.POST("/logout",handlers.LogoutHandler())
	r.POST("/register",handlers.RegisterHandler(database))
	r.POST("/login",handlers.LoginHandler(database))
	r.POST("/admin/create/",handlers.AdminCreateUser(database))
	r.POST("/admin/edit/:id/:role",handlers.AdminEditUser(database))
	r.POST("/admin/delete/:id",handlers.AdminDeleteUser(database))
	r.POST("/search-name",handlers.AdminUserSearch(database))

	

	r.Run(":8080")
}