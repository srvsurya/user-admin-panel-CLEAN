package handlers

import (
"github.com/gin-gonic/gin"
"golang.org/x/crypto/bcrypt"
"gorm.io/gorm"
"Week_12/models"
)

func AdminCreateUser(db *gorm.DB)gin.HandlerFunc{return func(c *gin.Context){
	var user models.User
	var role string
	name:=c.PostForm("name")
	email:=c.PostForm("email")
	create_role:=c.PostForm("role")
	password:=c.PostForm("password")
	if create_role=="user"{
		role="app_user"
	}else{
		role="admin"
	}
	hashpassword,err:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)

	user = models.User{
		Name:name,
		Email:email,
		Role:role,
		Password:string(hashpassword),
	}

	err = db.Create(&user).Error
	if err!=nil{
		
		c.JSON(500,gin.H{
			"message":"DB error"})
		return
	}
	c.Redirect(302,"/adminpanel")


}}