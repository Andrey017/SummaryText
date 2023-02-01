package db

import (
	"article_service/pkg/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func Init(url string) Database {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Article{})
	db.AutoMigrate(&models.Summary{})

	return Database{db}
}
