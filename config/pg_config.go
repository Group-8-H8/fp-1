package config

import (
	"fmt"
	// "log"
	"os"

	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func init() {
// 	if err := godotenv.Load(); err != nil {
// 		log.Fatalln(err.Error())
// 	}
// }

func GetDBConfig() gorm.Dialector {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	return postgres.Open(dsn)
}
