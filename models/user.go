package models

type User struct{
	UserID int `gorm:"column:user_id;primaryKey"`
	Name string
	Email string
	Role string
	Password string

}