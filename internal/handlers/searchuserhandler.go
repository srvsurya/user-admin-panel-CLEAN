package handlers
// Logging is usually unecessary here maybe?
import (
"github.com/gin-gonic/gin"
"Week_12/internal/models"
"Week_12/internal/service"
"Week_12/internal/logger"
"github.com/gin-contrib/sessions"
"go.uber.org/zap")

func AdminUserSearch(userService *service.UserService)gin.HandlerFunc{return func(c *gin.Context){
	session := sessions.Default(c)
	requestID := session.Get("request_id")
	id := session.Get("user_id")
	idInt,ok:=id.(int)
	if !ok{
		logger.Log.Error("Type assertion error",zap.Any("request_id",requestID))
	}
	search:=c.PostForm("search")
	var users []models.User

	users,err := userService.SearchUsers(search)
	if err!=nil{
		logger.Log.Error("Error at User Search",zap.Int("Admin_id",idInt),zap.Any("request_id",requestID))
		c.JSON(500,gin.H{
			"message":"DB error",
		})
		return
	}
	c.HTML(200,"admin_panel.html",gin.H{"users":users,})
}}