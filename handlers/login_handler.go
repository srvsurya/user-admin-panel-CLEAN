package handlers

import (
"github.com/gin-gonic/gin"
"golang.org/x/crypto/bcrypt"
"github.com/gin-contrib/sessions"
"Week_12/models"
"gorm.io/gorm"
"Week_12/logger"
"go.uber.org/zap"
)

func LoginHandler(db *gorm.DB)gin.HandlerFunc{return func(c* gin.Context){
	var pw_flag bool

	var user models.User


	email:=c.PostForm("email")
	password:=c.PostForm("password")
	requestID,_:=c.Get("request_id")
	logger.Log.Info("Login handler has started",zap.String("email",email),zap.Any("request_id",requestID))

	err := db.Select("user_id,name,email,password,role").Where("email = ?",email).First(&user).Error
	if err!=nil{
		if err==gorm.ErrRecordNotFound{
			logger.Log.Warn("Email does not exist in the db",zap.String("email",email),zap.Any("request_id",requestID))
			c.JSON(401,gin.H{"message":"Invalid credentials"})
			return
		}else{
			logger.Log.Error("Database query error",zap.String("email",email),zap.Error(err),zap.Any("request_id",requestID))
			c.JSON(500, gin.H{"message":"Server error"})
			return
		}
		
	}
	err=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err!=nil{
		pw_flag=false
	}else{
		pw_flag=true
	}
	if pw_flag==true{
		session:=sessions.Default(c)
		session.Set("user_id",user.UserID)
		session.Set("role",user.Role)
		err=session.Save()
		if err!=nil{
				logger.Log.Error("Session save failure",zap.Error(err),zap.Any("request_id",requestID))
				c.JSON(500,gin.H{"message":"Internal server error"})
				return
		}

		if user.Role=="admin"{
			logger.Log.Info("Admin login successful",zap.String("email",email),zap.Any("request_id",requestID))
			c.Redirect(302,"/adminpanel")
		} else{
			logger.Log.Info("User login successful",zap.String("email",email),zap.Any("request_id",requestID))
			c.Redirect(302,"/home")
		}
	}else{
		c.JSON(401,gin.H{"message":"Invalid credentials"})
		logger.Log.Warn("Password mismatch",zap.String("email",email),zap.Any("request_id",requestID))
		return

	}

	
	
}

}