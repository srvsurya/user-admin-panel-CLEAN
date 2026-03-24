package handlers

import (
	
	"github.com/gin-gonic/gin"
	"Week_12/internal/logger"
	
	"go.uber.org/zap"
	"Week_12/internal/service"
)

func RegisterHandler(userService *service.UserService)gin.HandlerFunc {return func(c *gin.Context){
	
	name:=c.PostForm("name")
	email:=c.PostForm("email")
	password:=c.PostForm("password")
	confirmPassword:=c.PostForm("confirmpassword")
	requestID,_:=c.Get("request_id")
	logger.Log.Info("Register handler has started",zap.String("email",email),zap.Any("request_id",requestID))
	
	err := userService.RegisterUser(name,email,password,confirmPassword)
	if err!=nil{
		logger.Log.Error("Error at User registration",zap.String("email",email),zap.Any("request_id",requestID),zap.Error(err))
		c.JSON(500,gin.H{"message":"Server side error"})
		return
	}

	logger.Log.Info("User account created successfully",zap.String("email",email),zap.Any("request_id",requestID))
	c.Redirect(302,"/login")


}
}