package database

import (
	"crud-go/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDatabase() *gorm.DB {
	dsn:="host=localhost user=postgres password=123456 dbname=crud-go port=5432 sslmode=disable "
	db,err:= gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connect database success")
	db.AutoMigrate(&model.User{},&model.Product{})
	fmt.Println("database migrated")
	DB = db
	return db
}