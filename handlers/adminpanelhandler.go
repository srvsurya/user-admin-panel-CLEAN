package handlers

import (
"github.com/gin-gonic/gin"
"Week_12/models"
"strconv"
"gorm.io/gorm"

)

func AdminviewHandler(db *gorm.DB)gin.HandlerFunc{
	return func(c* gin.Context){
		limit:=7
		page_string:=c.DefaultQuery("page","1")
		page,err:=strconv.Atoi(page_string)
		if err!=nil || page<1{
			page=1
		}
		offset:=(page-1)*limit

		var users []models.User
		
		err = db.Order("user_id").
		Limit(limit).
		Offset(offset).
		Find(&users).Error
		if err!=nil{
				c.JSON(500,gin.H{"message":"Database error"})
				return
		}
		c.HTML(200,"admin_panel.html",gin.H{"users": users,"page":page,})
	}
}