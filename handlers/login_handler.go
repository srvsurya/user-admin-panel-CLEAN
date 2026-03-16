package handlers

import (
"github.com/gin-gonic/gin"
"golang.org/x/crypto/bcrypt"
"github.com/gin-contrib/sessions"
"Week_12/models"
"gorm.io/gorm"
)

func LoginHandler(db *gorm.DB)gin.HandlerFunc{return func(c* gin.Context){
	var pw_flag bool

	var user models.User


	email:=c.PostForm("email")
	password:=c.PostForm("password")

	err := db.Select("name,email,password,role").Where("email = ?",email).First(&user).Error
	if err!=nil{
		if err==gorm.ErrRecordNotFound{
			c.JSON(401,gin.H{"message":"Invalid credentials"})
			return
		}else{
			c.JSON(500, gin.H{"message":"Server error"})
			return
		}
		
	}
	err=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
	if err!=nil{
		pw_flag=false
	}else{
		pw_flag=true
	}
	if pw_flag==true{
		session:=sessions.Default(c)
		session.Set("user_id",user.UserID)
		session.Set("role",user.Role)
		session.Save()

		if user.Role=="admin"{
			c.Redirect(302,"/adminpanel")
		} else{
			c.Redirect(302,"/home")
		}
	}else{
		c.JSON(401,gin.H{"message":"Invalid credentials"})
		return

	}

	
	
}

}