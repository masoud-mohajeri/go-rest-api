package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}

type Articles []Article

func GetAllArticles(c *gin.Context) {
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
	c.JSON(http.StatusOK, articles)
}
