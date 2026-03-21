package handlers

import (
	"Week_12/logger"
	"Week_12/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func AdminDeleteUser(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		id := c.Param("id")
		requestID,_:=c.Get("request_id")
		idInt, err := strconv.Atoi(id)

		if err!= nil {
			logger.Log.Error("Conversion error", zap.String("user_id", id), zap.Error(err),zap.Any("request_id",requestID))
			return
		}
		logger.Log.Info("Admin delete handler has started", zap.Int("user_id", idInt),zap.Any("request_id",requestID))

		err = db.Delete(&models.User{}, idInt).Error
		if err != nil {
			logger.Log.Error("Database deletion error", zap.Int("user_id", idInt), zap.Error(err),zap.Any("request_id",requestID))
			c.JSON(500, gin.H{"message": "DB error"})
			return
		}
		logger.Log.Info("User deleted", zap.Int("user_id", idInt),zap.Any("request_id",requestID))
		c.Redirect(302, "/adminpanel")
	}
}
