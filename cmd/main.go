package main

import (
	"Week_12/internal/db"
	"Week_12/internal/handlers"
	"Week_12/internal/middleware"
	"net/http"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"Week_12/internal/logger"
	"Week_12/internal/service"
	"Week_12/internal/repository"

)

func main(){
	logger.Init()
	defer logger.Log.Sync()
	database:=db.Connect()
	db.Migrate(database)
	r:=gin.Default()
	userRepo:=repository.NewUserRepository(database)
	userService:=service.NewUserService(userRepo)
	//template functions
	r.SetFuncMap(template.FuncMap{
		"add":func( a int, b int )int{return a+b},
		"sub":func(a int, b int)int{return a-b},
	})
	r.LoadHTMLGlob("web/templates/*")
	r.Static("web/static","./static")

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
	r.GET("/home",middleware.UserMiddlewareAuth(),handlers.HomeHandler(userService))
	r.GET("/login",func(c *gin.Context){
		c.HTML(200,"login.html",nil)
	})
	r.GET("/adminpanel",middleware.UserMiddlewareAuth(),middleware.AdminMiddlewareAuth(),handlers.AdminviewHandler(userService))
	r.GET("/admin/create",middleware.UserMiddlewareAuth(),middleware.AdminMiddlewareAuth(),func(c *gin.Context){
		c.HTML(200,"create_user.html",nil)
	})
	r.POST("/logout",handlers.LogoutHandler())
	r.POST("/register",handlers.RegisterHandler(userService))
	r.POST("/login",handlers.LoginHandler(userService))
	r.POST("/admin/create/",handlers.AdminCreateUser(userService))
	r.POST("/admin/edit/:id/:role",handlers.AdminEditUser(userService))
	r.POST("/admin/delete/:id",handlers.AdminDeleteUser(userService))
	r.POST("/search-name",handlers.AdminUserSearch(userService))

	

	r.Run(":8080")
}