package db

import (
	_ "github.com/lib/pq"
		"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"Week_12/internal/models"
	"fmt"
	"Week_12/internal/config"
	"os"

)


func Connect()*gorm.DB{
	config.LoadEnv()
	dsn:=fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PORT"))
	db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Fatal(err)
	}
	
	return db

}
func Migrate(db *gorm.DB){
	err := db.AutoMigrate(&models.User{})
	if err!=nil{
		log.Fatalf("Failed to automigrate database: %v",err)
	}
}

