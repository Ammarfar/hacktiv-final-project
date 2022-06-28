package configs

import (
	"finalproject/helpers"
	"finalproject/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = helpers.GetEnv("DB_HOST")
	user     = helpers.GetEnv("DB_USER")
	pass     = helpers.GetEnv("DB_PASS")
	name     = helpers.GetEnv("DB_NAME")
	port     = helpers.GetEnv("DB_PORT")
	timeZone = helpers.GetEnv("DB_TIMEZONE")
	db       *gorm.DB
	err      error
)

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=%s", host, user, pass, name, port, timeZone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.Photo{},
		&models.Comment{},
		&models.SocialMedia{},
	)
}

func GetDB() *gorm.DB {
	return db
}
