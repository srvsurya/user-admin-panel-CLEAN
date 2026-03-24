package handlers

import (
	"Week_12/internal/logger"
	"Week_12/internal/models"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"Week_12/internal/service"
)

func AdminviewHandler(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		idInterface := session.Get("user_id")
		requestID, _ := c.Get("request_id")
		id, ok := idInterface.(int)
		page_string := c.DefaultQuery("page", "1")
		limit := 7
		page, err := strconv.Atoi(page_string)
		if err != nil || page < 1 {
			page = 1
		}
		offset := (page - 1) * limit

		
		
		
		if !ok {
			logger.Log.Error("Type assertion error at sessions", zap.Any("request_id", requestID))
			c.JSON(500,gin.H{"message":"Internal Server Error"})
			return
		}
		logger.Log.Info("admin panel handler has started", zap.Int("user_id", id), zap.Any("request_id", requestID))

		var users []models.User
		users,err = userService.AdminViewService(id,offset,limit)
		if err!=nil{
			logger.Log.Error("User View extraction failure",zap.Int("user_id",id),zap.Any("request_id",requestID),zap.Error(err))
			c.JSON(500,gin.H{"message":"Internal Server Error"})
			return
		}
		
		logger.Log.Info("Admin panel view successfully loaded", zap.Int("user_id", id), zap.Any("request_id", requestID))
		c.HTML(200, "admin_panel.html", gin.H{"users": users, "page": page})
	}
}
