package handlers

import (
"github.com/gin-gonic/gin"
"gorm.io/gorm"
"Week_12/models"
"Week_12/logger"
"go.uber.org/zap"
"github.com/gin-contrib/sessions"
)

func AdminCreateUser(db *gorm.DB)gin.HandlerFunc{return func(c *gin.Context){
	var user models.User
	var role string
	session:=sessions.Default(c)
	idInterface:=session.Get("user_id")
	requestID,_:=c.Get("request_id")
	id,ok:=idInterface.(int)
	if !ok{
		logger.Log.Info("Type assertion error")
		return
	}
	logger.Log.Info("Create User at Admin side handler has started",zap.Int("admin_id",id),zap.Any("request_id",requestID))
	name:=c.PostForm("name")
	email:=c.PostForm("email")
	create_role:=c.PostForm("role")
	password:=c.PostForm("password")
	if create_role=="user"{
		role="app_user"
	}else{
		role="admin"
	}
	


	user = models.User{
		Name:name,
		Email:email,
		Role:role,
		Password:password,
	}

	err := db.Create(&user).Error
	if err!=nil{
		logger.Log.Error("Creation at database failed",zap.Int("admin_id",id),zap.String("user_name",name),zap.Error(err),zap.Any("request_id",requestID))
		c.JSON(500,gin.H{
			"message":"DB error"})
		return
	}
	logger.Log.Info("User created at admin side",zap.Int("admin_id",id),zap.String("user_name",name),zap.Any("request_id",requestID))
	c.Redirect(302,"/adminpanel")


}}