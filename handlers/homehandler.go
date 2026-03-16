package handlers

import (
	"Week_12/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HomeHandler(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		id := session.Get("user_id")

		var user models.User

		err := db.Select("name").Where("user_id = ?", id).First(&user).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(404, gin.H{"message": "User does not exist"})

			} else {
				c.JSON(500, gin.H{"message": "DB error"})
			}
			return

		}
		c.HTML(200, "home.html", gin.H{"name": user.Name})
	}
}
