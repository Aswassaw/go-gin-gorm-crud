package config

import (
	"fmt"
	"go-gin-gorm-crud/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	user     = "root"
	password = ""
	host     = "localhost"
	port     = "3306"
	dbName   = "go-simple-crud"
)

func DatabaseConnection() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)

	return db
}
