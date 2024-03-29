package providers

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() {
	conn, err := gorm.Open(sqlserver.Open("server=host.docker.internal;user id=sa;password=@Un1v3r53C4t4l0g;port=1401"), &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("An error occurred while trying to create database connection: %s", err.Error()))
	}

	db = conn
}

func GetDB() *gorm.DB {
	return db
}
