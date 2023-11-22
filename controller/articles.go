package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}

type Articles []Article

func GetAllArticles(rw http.ResponseWriter, req *http.Request) {
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
