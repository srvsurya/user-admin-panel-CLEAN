package handlers

import (
	"Week_12/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"Week_12/logger"
	"go.uber.org/zap"
)

func HomeHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		idInterface := session.Get("user_id")
		requestID,_:=c.Get("request_id")
		id,ok:=idInterface.(int)
		if !ok{
			logger.Log.Info("Failed at type assertion for session id",zap.Any("request_id",requestID))
		}
		logger.Log.Info("Home handler started",zap.Int("user_id",id),zap.Any("request_id",requestID))
		var user models.User

		err := db.Select("name").Where("user_id = ?", id).First(&user).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				logger.Log.Warn("user_id does not exist",zap.Int("user_id",id),zap.Any("request_id",requestID))
				c.JSON(404, gin.H{"message": "User does not exist"})

			} else {
				logger.Log.Error("Database Query Error",zap.Error(err),zap.Int("user_id",id),zap.Any("request_id",requestID))
				c.JSON(500, gin.H{"message": "DB error"})
			}
			return

		}
		logger.Log.Info("Home screen successfully loaded",zap.Int("user_id",id),zap.Any("request_id",requestID))
		c.HTML(200, "home.html", gin.H{"name": user.Name})
	}
}
