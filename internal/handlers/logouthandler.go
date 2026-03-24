package handlers

import (
	"Week_12/internal/logger"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LogoutHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		idInterface := session.Get("user_id")
		requestID, _ := c.Get("request_id")
		id, ok := idInterface.(int)
		if !ok {
			logger.Log.Error("Type Assertion failure", zap.Any("request_id", requestID))
		}
		logger.Log.Info("Logout handler has started", zap.Int("user_id", id), zap.Any("request_id", requestID))
		session.Clear()
		err := session.Save()
		if err != nil {
			logger.Log.Error("Session save failure", zap.Int("user_id", id), zap.Error(err), zap.Any("request_id", requestID))
			c.JSON(500, gin.H{"message": "Please logout again"})
		}

		logger.Log.Info("User logged out", zap.Int("user_id", id), zap.Any("request_id", requestID))

		c.Redirect(302, "/login")
	}
}
