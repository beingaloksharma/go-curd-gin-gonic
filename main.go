package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Article struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Author      *Author `json:"author"`
}

type Author struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

var article []Article

func createArticle(c *gin.Context) {
	var art Article
	RequestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid Parameter"})
		return
	}
	json.Unmarshal(RequestBody, &art)
	article = append(article, art)
	c.JSON(http.StatusCreated, gin.H{
		// "data":    article,
		"message": "article saved successfully",
	})
	return
}

func getAllArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": article,
	})
	return
}

func deleteArticle(c *gin.Context) {
	id := c.Param("id")
	_deleteArticleByid(id)
	c.JSON(http.StatusOK, gin.H{
		"data": article,
	})
	return
}

func updateArticle(c *gin.Context) {
	var art Article
	RequestBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "Invalid Parameter"})
		return
	}
	json.Unmarshal(RequestBody, &art)
	id := c.Param("id")
	art.ID = id
	_deleteArticleByid(id)
	article = append(article, art)
	c.JSON(http.StatusOK, gin.H{
		"data":    article,
		"message": "article updated successfully",
	})
	return
}

func _deleteArticleByid(id string) {
	for index, art := range article {
		if art.ID == id {
			article = append(article[:index], article[index+1:]...)
			break
		}
	}
}

func getArticleById(c *gin.Context) {
	id := c.Param("id")
	for _, art := range article {
		if art.ID == id {
			c.JSON(http.StatusCreated, gin.H{
				"data": &art,
			})
			return
		}
	}
}

func handleRequest() {
	route := gin.Default()
	route.POST("/create-article", createArticle)
	route.GET("/", getAllArticle)
	route.GET("/get-all-article", getAllArticle)
	route.GET("/get-by-article/:id", getArticleById)
	route.DELETE("/delete-article-by-id/:id", deleteArticle)
	route.PUT("/update-article-by-id/:id", updateArticle)
	route.Run(":8080")
}

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	handleRequest()
}
