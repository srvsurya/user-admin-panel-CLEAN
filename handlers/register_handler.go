package handlers

import (
	
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"Week_12/models"
	"Week_12/logger"
	"go.uber.org/zap"
)

func RegisterHandler(db *gorm.DB)gin.HandlerFunc {return func(c *gin.Context){
	
	name:=c.PostForm("name")
	email:=c.PostForm("email")
	password:=c.PostForm("password")
	confirmPassword:=c.PostForm("confirmpassword")
	role:="app_user"
	requestID,_:=c.Get("request_id")
	var user models.User
	logger.Log.Info("Register handler has started",zap.String("email",email),zap.Any("request_id",requestID))
	
	

	if name=="" || email=="" || password=="" || confirmPassword==""{
		c.JSON(400,gin.H{"message":"Missing required fields"})
		return
	}
	if password!=confirmPassword{
		c.JSON(400,gin.H{"message":"Passwords don't match"})
		logger.Log.Warn("Password doesn't match ",zap.String("email",email),zap.Any("request_id",requestID))
		return
	}
	
	var count int64
	user = models.User{Name: name,
				Email: email,
				Password: password,
				Role: role}

	err:=db.Model(&models.User{}).Where("email = ?",email).Count(&count).Error
	if count>0{
		c.JSON(400,gin.H{"error":"Email already exists"})
		logger.Log.Warn("Email used for registeration already exists",zap.String("email",email),zap.Any("request_id",requestID))
		return
	}else if err!=nil{
		logger.Log.Error("Failed Query",zap.Error(err),zap.Any("request_id",requestID))
		c.JSON(500,gin.H{"message":"Internal DB error"})
		return
	}
	
	err = db.Create(&user).Error
	if err!=nil{
		c.JSON(500,gin.H{"message":"Failed to create user"})
		logger.Log.Error("Failed at table insert",zap.String("email",email),zap.Error(err),zap.Any("request_id",requestID))
		return
	}
	logger.Log.Info("User account created successfully",zap.String("email",email),zap.Any("request_id",requestID))
	c.Redirect(302,"/login")


}
}