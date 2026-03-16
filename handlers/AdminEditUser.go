package handlers

import (
"github.com/gin-gonic/gin"
"Week_12/models"
"gorm.io/gorm"
)

func AdminEditUser(db *gorm.DB)gin.HandlerFunc{return func(c* gin.Context){
	var role string
	id:=c.Param("id")
	choice:=c.Param("role")

	



	if choice=="admin"{
		role="app_user"
	}else{
		role="admin"
	}

	err := db.Model(&models.User{}).Where("user_id = ?",id).Update("role",role).Error
	if err!=nil{
		c.JSON(500,gin.H{"message":"DB error"})
		return
	}
	c.Redirect(302,"/adminpanel")
}}