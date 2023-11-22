package main

import (
	"log"
	"net/http"

	"github.com/mohajerimasoud/go-rest-api/controller"
)

func appRouter() {
	http.HandleFunc("/articles", controller.GetAllArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	appRouter()
}
