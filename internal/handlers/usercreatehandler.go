package handlers

import (
	"Week_12/internal/logger"

	"Week_12/internal/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func AdminCreateUser(userService *service.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {

		var role string
		session := sessions.Default(c)
		idInterface := session.Get("user_id")
		requestID, _ := c.Get("request_id")
		id, ok := idInterface.(int)
		if !ok {
			logger.Log.Info("Type assertion error")
			return
		}
		logger.Log.Info("Create User at Admin side handler has started", zap.Int("admin_id", id), zap.Any("request_id", requestID))
		name := c.PostForm("name")
		email := c.PostForm("email")
		create_role := c.PostForm("role")
		password := c.PostForm("password")
		if create_role == "user" {
			role = "app_user"
		} else {
			role = "admin"
		}
		err := userService.AdminRegisterUser(name, email, role, password)
		if err != nil {
			logger.Log.Error("Error at Admin user creation", zap.Int("user_id", id), zap.Any("request_id", requestID), zap.Error(err))
			c.JSON(500, gin.H{"message": "Internal Server Error"})
			return
		}
		logger.Log.Info("User created at admin side", zap.Int("admin_id", id), zap.String("user_name", name), zap.Any("request_id", requestID))
		c.Redirect(302, "/adminpanel")

	}
}
