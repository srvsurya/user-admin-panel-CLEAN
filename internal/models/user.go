package models

import ("Week_12/internal/utils"
"gorm.io/gorm"
"Week_12/internal/logger"
"go.uber.org/zap")

type User struct{
	UserID int `gorm:"column:user_id;primaryKey"`
	Name string
	Email string`gorm:"unique"`
	Role string
	Password string

}

func (u *User) BeforeCreate(tx *gorm.DB)(err error){
	
	u.Password,err = utils.HashPassword(u.Password,u.Email)
	if err!=nil{
		logger.Log.Error("Hook failure",zap.String("email",u.Email),zap.Error(err))
	}
	return
}