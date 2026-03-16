package handlers

import (
	
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"fmt"
	"gorm.io/gorm"
	"Week_12/models"
)

func RegisterHandler(db *gorm.DB)gin.HandlerFunc {return func(c *gin.Context){
	name:=c.PostForm("name")
	email:=c.PostForm("email")
	password:=c.PostForm("password")
	confirmPassword:=c.PostForm("confirmpassword")
	role:="app_user"
	var user models.User
	
	

	if name=="" || email=="" || password=="" || confirmPassword==""{
		c.JSON(400,gin.H{"message":"Missing required fields"})
		return
	}
	if password!=confirmPassword{
		c.JSON(400,gin.H{"message":"Passwords don't match"})
		return
	}
	hashpassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	if err!=nil{
		c.JSON(500,gin.H{"message":"Server error"})
		return
	}
	var count int64
	user = models.User{Name: name,
				Email: email,
				Password: string(hashpassword),
				Role: role}

	err=db.Model(&models.User{}).Where("email = ?",email).Count(&count).Error
	if count>0{
		c.JSON(400,gin.H{"error":"Email already exists"})
		return
	}else if err!=nil{
		c.JSON(500,gin.H{"message":"Internal DB error"})
		return
	}
	
	err = db.Create(&user).Error
	if err!=nil{
		fmt.Println("Error:",err)
		c.JSON(500,gin.H{"message":"Failed to create user"})
		return
	}
	c.Redirect(302,"/login")


}
}