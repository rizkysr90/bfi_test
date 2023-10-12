package database

import (
	"myapp/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
func Connection() *gorm.DB{
	dsn := "host=localhost user=postgres password=adarizki123 dbname=bfi_test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db,_ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&model.Ticket{})
	db.AutoMigrate(&model.User{})
	return db
}