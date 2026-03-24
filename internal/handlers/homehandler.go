package handlers

import (
	"Week_12/internal/service"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	
	"Week_12/internal/logger"
	"go.uber.org/zap"
)

func HomeHandler(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		idInterface := session.Get("user_id")
		requestID,_:=c.Get("request_id")
		id,ok:=idInterface.(int)
		if !ok{
			logger.Log.Info("Failed at type assertion for session id",zap.Any("request_id",requestID))
		}
		logger.Log.Info("Home handler started",zap.Int("user_id",id),zap.Any("request_id",requestID))
		

		
		name,err:=userService.HomeService(id)
		if err!=nil{
			logger.Log.Error("Error at Home page name extraction",zap.Any("request_id",requestID),zap.Int("user_id",id),zap.Error(err))
			c.JSON(500,gin.H{"message":"Internal Server Error"})
			return
		}
		logger.Log.Info("Home screen successfully loaded",zap.Int("user_id",id),zap.Any("request_id",requestID))
		c.HTML(200, "home.html", gin.H{"name": name})
	}
}
