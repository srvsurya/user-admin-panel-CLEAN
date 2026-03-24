package handlers

import (
	"Week_12/internal/logger"
	

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"Week_12/internal/service"
)

func LoginHandler(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		

		

		email := c.PostForm("email")
		password := c.PostForm("password")
		requestID, _ := c.Get("request_id")
		logger.Log.Info("Login handler has started", zap.String("email", email), zap.Any("request_id", requestID))

		user,err:=userService.Login(email,password)
		if err!=nil{
			logger.Log.Error("Invalid Credentials detected",zap.String("email",email),zap.Any("request_id",requestID),zap.Error(err))
			c.JSON(401,gin.H{"message":"Invalid Credentials"})
			return
		}
		
		session := sessions.Default(c)
		session.Set("user_id", user.UserID)
		session.Set("role", user.Role)
		err = session.Save()
		if err != nil {
			logger.Log.Error("Session save failure", zap.Error(err), zap.Any("request_id", requestID))
			c.JSON(500, gin.H{"message": "Internal server error"})
			return
		}

		if user.Role == "admin" {
			logger.Log.Info("Admin login successful", zap.String("email", email), zap.Any("request_id", requestID))
			c.Redirect(302, "/adminpanel")
		} else {
			logger.Log.Info("User login successful", zap.String("email", email), zap.Any("request_id", requestID))
			c.Redirect(302, "/home")
		}

	}

}
