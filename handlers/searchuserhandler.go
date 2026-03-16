package handlers

import (
"github.com/gin-gonic/gin"
"gorm.io/gorm"
"Week_12/models")

func AdminUserSearch(db *gorm.DB)gin.HandlerFunc{return func(c *gin.Context){
	search:=c.PostForm("search")
	var users []models.User

	err := db.Select("user_id,name,email,password,role").Where("name ILIKE ?", "%"+search+"%").Find(&users).Error
	if err!=nil{
		c.JSON(500,gin.H{
			"message":"DB error",
		})
		return
	}
	c.HTML(200,"admin_panel.html",gin.H{"users":users,})
}}