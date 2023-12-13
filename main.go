package main

import (
	"github.com/mohajerimasoud/go-rest-api/config"
	"github.com/mohajerimasoud/go-rest-api/database"
	"github.com/mohajerimasoud/go-rest-api/models"
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	// TODO: handle error for all migrations with a good pattern
	db.AutoMigrate(&models.Article{})
}

func main() {
	// TODO: read values from env
	db, err := database.ConnectToDB("pg", "pass", "crud")

	if err != nil {
		// TODO: add logger
		log.Fatalln(err)
	}

	Migrate(db)

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	route := config.SetupRoutes(db)

	log.Fatalln(route.Run(":8000"))
}
