package handlers

import ("gorm.io/gorm"
"github.com/gin-gonic/gin"
"Week_12/models"
)

func AdminDeleteUser(db *gorm.DB)gin.HandlerFunc{return func(c *gin.Context){
	id:=c.Param("id")
	err := db.Delete(&models.User{},id).Error
	if err!=nil{
		c.JSON(500,gin.H{"message":"DB error"})
		return
	}
	c.Redirect(302,"/adminpanel")
}}