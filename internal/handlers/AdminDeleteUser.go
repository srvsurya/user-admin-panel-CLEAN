package handlers

import (
	"Week_12/internal/logger"
	"strconv"
	"Week_12/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/gin-contrib/sessions"
)

func AdminDeleteUser(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		session:=sessions.Default(c)
		adminId:=session.Get("user_id")
		requestID,_:=c.Get("request_id")
		adminIdInt,ok:=adminId.(int)
		if !ok{
			logger.Log.Error("Type assertion error",zap.Any("request_id",requestID))
			c.JSON(500,gin.H{"message":"Internal Server Error"})
			return
		}
		id := c.Param("id")
		
		idInt, err := strconv.Atoi(id)

		if err!= nil {
			logger.Log.Error("Conversion error", zap.String("user_id", id), zap.Error(err),zap.Any("request_id",requestID))
			return
		}
		logger.Log.Info("Admin delete handler has started", zap.Int("user_id", idInt),zap.Any("request_id",requestID))

		err=userService.DeleteUser(idInt,adminIdInt)
		if err!=nil{
			logger.Log.Error("Error at User Deletion in the admin side",zap.Int("user_id",idInt),zap.Any("request_id",requestID),zap.Error(err))
			c.JSON(500,gin.H{"message":"Internal Server Error"})
			return
		}
		logger.Log.Info("User deleted", zap.Int("user_id", idInt),zap.Any("request_id",requestID))
		c.Redirect(302, "/adminpanel")
	}
}
