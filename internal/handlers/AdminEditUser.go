package handlers

import (
	"Week_12/internal/logger"

	"strconv"
	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"Week_12/internal/service"
)

func AdminEditUser(userService *service.UserService)gin.HandlerFunc{return func(c* gin.Context){
	
	id:=c.Param("id")
	session:=sessions.Default(c)
	
	requestID,_:=c.Get("request_id")
	adminId:=session.Get("user_id")
	adminIdInt,ok:=adminId.(int)
	if !ok{
		logger.Log.Error("Type assertion error",zap.Any("request_id",requestID))
		c.JSON(500,gin.H{"message":"Internal Server Error"})
		return
	}
	current_role:=c.Param("role")
	idInt,err:=strconv.Atoi(id)
	if err!=nil{
		logger.Log.Error("Conversion error",zap.Int("user_id",idInt),zap.Error(err),zap.Any("request_id",requestID))
		return
	}
	logger.Log.Info("CRUD Edit handler has started",zap.Int("user_id",int(idInt)),zap.Any("request_id",requestID))

	err=userService.AdminEdit(idInt,adminIdInt,current_role)
	if err!=nil{
		logger.Log.Error("Error on Admin Edit",zap.Int("user_id",idInt),zap.Any("request_id",requestID),zap.Error(err))
		c.JSON(500,gin.H{"message":"Internal Server Error"})
		return
	}

	
	logger.Log.Info("Role has been updated successfully",zap.Int("user_id",idInt),zap.Any("request_id",requestID))
	c.Redirect(302,"/adminpanel")
}}