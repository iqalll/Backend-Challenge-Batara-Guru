package db

import (
	"Backend-Challenge-Batara-Guru/db/seeders"
	"Backend-Challenge-Batara-Guru/models"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbGorm *gorm.DB

func NewGorm() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConnect := os.Getenv("DB_DSN")
	dbGorm, err = gorm.Open(mysql.Open(dbConnect), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	boolValue, _ := strconv.ParseBool(os.Getenv("IS_AUTOMIGRATE"))
	if boolValue {
		dbGorm.AutoMigrate(
			&models.User{},
			&models.Gift{},
		)
		seeders.DBSeed(dbGorm)

	}

	return dbGorm
}
