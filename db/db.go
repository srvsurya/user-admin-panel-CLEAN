package db

import (
	_ "github.com/lib/pq"
		"log"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"

)

func Connect()*gorm.DB{
	dsn:="host=localhost user=saurav password=admin dbname=db_users port=5432 sslmode=disable"
	db,err:=gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err!=nil{
		log.Fatal(err)
	}
	return db

}

