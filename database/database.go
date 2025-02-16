package database

import (
	"fmt"
	"log"

	"avito-assignment-2025/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(conf *config.Config) {
	conn_str := fmt.Sprintf(
		"host=%s port=%s user=%s "+
			"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.DBPort, conf.User, conf.Password, conf.Name,
	)

	conn, err := gorm.Open(postgres.Open(conn_str), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect DB: ", err)
	}
	DB = conn

	log.Println("DB succesfully connected")
}
