package handlers

import (
"github.com/gin-gonic/gin"
"Week_12/models"
"gorm.io/gorm"
"Week_12/logger"
"go.uber.org/zap"
"strconv"

)

func AdminEditUser(db *gorm.DB)gin.HandlerFunc{return func(c* gin.Context){
	var role string
	id:=c.Param("id")
	requestID,_:=c.Get("request_id")
	choice:=c.Param("role")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		logger.Log.Error("Conversion error",zap.Int("user_id",idInt),zap.Error(err),zap.Any("request_id",requestID))
		return
	}
	logger.Log.Info("CRUD Edit handler has started",zap.Int("user_id",int(idInt)),zap.Any("request_id",requestID))

	if choice=="admin"{
		role="app_user"
	}else{
		role="admin"
	}

	err = db.Model(&models.User{}).Where("user_id = ?",id).Update("role",role).Error
	if err!=nil{
		logger.Log.Error("Database updation error",zap.Error(err),zap.Int("user_id",idInt),zap.Any("request_id",requestID))
		c.JSON(500,gin.H{"message":"DB error"})
		return
	}
	logger.Log.Info("Role has been updated successfully",zap.Int("user_id",idInt),zap.Any("request_id",requestID))
	c.Redirect(302,"/adminpanel")
}}