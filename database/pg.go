package database

import (
	"log"

	"github.com/Group-8-H8/fp-1/config"
	"github.com/Group-8-H8/fp-1/entity"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	db, err = gorm.Open(config.GetDBConfig())
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err = db.AutoMigrate(entity.Todo{}); err != nil {
		log.Fatalln(err.Error())
	}

	log.Println("database connected!")
}

func GetDbInstance() *gorm.DB {
	return db
}
