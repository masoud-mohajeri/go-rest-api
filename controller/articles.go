package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
	Id      int16  `json:"Id"`
}

type Articles []Article

func GetAllArticles(c *gin.Context) {
	articles := Articles{
		Article{
			Id:      1,
			Title:   "A1",
			Desc:    "D1",
			Content: "Hello world",
		},
		Article{
			Id:      2,
			Title:   "A2",
			Desc:    "D2",
			Content: "Hello Guys",
		},
	}

	c.JSON(http.StatusOK, articles)
}

func CreateArticle(c *gin.Context) {
	var article Article

	if err := c.BindJSON(&article); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// add user to db
	c.JSON(http.StatusCreated, article)
}
