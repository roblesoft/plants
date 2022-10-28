package db

import (
	"log"

	"github.com/roblesoft/plants/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Plant{})
	db.AutoMigrate(&models.Garden{})

	return db
}

// type Database struct {
// 	db *gorm.DB
// }

// var DB *gorm.DB

// func Init(url string) {
// 	if DB != nil {
// 		return
// 	}

// 	var err error
// 	database, err := gorm.Open(postgres.Open(url), &gorm.Config{})

// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	database.AutoMigrate(&models.Plant{})
// 	database.AutoMigrate(&models.Garden{})

// 	DB = database
// }
