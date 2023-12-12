package main

import (
	"github.com/mohajerimasoud/go-rest-api/config"
	"github.com/mohajerimasoud/go-rest-api/database"
	"github.com/mohajerimasoud/go-rest-api/models"
	"github.com/mohajerimasoud/go-rest-api/repositories"
	"log"
)

func main() {
	// TODO: read values from env
	db, err := database.ConnectToDB("pg", "pass", "crud")

	if err != nil {
		// TODO: add logger
		log.Fatalln(err)
	}

	// TODO: Migrate data
	migrationError := db.AutoMigrate(&models.Article{})

	if migrationError != nil {
		// TODO: add logger
		log.Fatalln(err)
	}
	//TODO: Unresolved reference 'Close' - find alternative
	//defer db.Close()

	contactRepository := repositories.NewArticleRepository(db)

	route := config.SetupRoutes(contactRepository)

	log.Fatalln(route.Run(":8000"))
}
