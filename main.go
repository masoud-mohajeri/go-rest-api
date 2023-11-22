package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}

type Articles []Article

func homePage(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello world!")
}

func allArticles(rw http.ResponseWriter, req *http.Request) {
	articles := Articles{
		Article{
			Title:   "A1",
			Desc:    "D1",
			Content: "Hello world",
		},
		Article{
			Title:   "A2",
			Desc:    "D2",
			Content: "Hello Guys",
		},
	}
	fmt.Println("All articles endpoint")
	json.NewEncoder(rw).Encode(articles)
}

func appRouter() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	appRouter()

}
