package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db *gorm.DB
)

func Connect() error {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_todo_app?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Println("Failed to connect to the database", err)
		return err
	}

	db = d
	log.Println("Database connected successfully")
	return nil
}

func GetDb() *gorm.DB {
	return db
}
