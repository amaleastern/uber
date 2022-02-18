package db

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"uber/src/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Init Database initialization
func Init() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=%s&parseTime=True&loc=%s",
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		"utf8",
		url.QueryEscape(os.Getenv("TIMEZONE")),
	)

	dbHandler, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               dsn,
		DefaultStringSize: 256,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	return dbHandler
}

// Migrate Automatically migrate your schema
func Migrate(dbHandler *gorm.DB) {
	dbHandler.AutoMigrate(&models.User{})
	dbHandler.AutoMigrate(&models.Booking{})
	dbHandler.AutoMigrate(&models.Cab{})
	dbHandler.AutoMigrate(&models.Location{})
}
