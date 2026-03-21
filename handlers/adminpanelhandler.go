package handlers

import (
"github.com/gin-gonic/gin"
"Week_12/models"
"strconv"
"gorm.io/gorm"
"Week_12/logger"
"go.uber.org/zap"
"github.com/gin-contrib/sessions"

)

func AdminviewHandler(db *gorm.DB)gin.HandlerFunc{
	return func(c* gin.Context){
		session:=sessions.Default(c)
		idInterface:=session.Get("user_id")
		requestID,_:=c.Get("request_id")
		id,ok:=idInterface.(int)
		if !ok{
			logger.Log.Error("Type assertion error at sessions",zap.Any("request_id",requestID))
			return
		}

		limit:=7
		page_string:=c.DefaultQuery("page","1")
		page,err:=strconv.Atoi(page_string)
		if err!=nil || page<1{
			page=1
		}
		offset:=(page-1)*limit

		var users []models.User
		logger.Log.Info("admin panel handler has started",zap.Int("user_id",id),zap.Any("request_id",requestID))
		
		err = db.Order("user_id").
		Limit(limit).
		Offset(offset).
		Find(&users).Error
		if err!=nil{
				logger.Log.Error("Database fetch error",zap.Error(err),zap.Int("user_id",id),zap.Any("request_id",requestID))
				c.JSON(500,gin.H{"message":"Database error"})
				return
		}
		logger.Log.Info("Admin panel view successfully loaded",zap.Int("user_id",id),zap.Any("request_id",requestID))
		c.HTML(200,"admin_panel.html",gin.H{"users": users,"page":page,})
	}
}