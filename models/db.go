package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// Migrating
	db = GetDB()

	log.Println("Migrating all models....")
	db.AutoMigrate(&User{}, &Class{})
	log.Println("Migrating success")
}

func GetDB() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/web-kafekoding?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}
