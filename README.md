# **Introduction**

  * Go, also known as Golang, is an open-source, compiled, and statically typed programming language designed by Google. It is built to be simple, high-performing, readable, and efficient.
  * Go supports concurrent programming, i.e. it allows running multiple processes simultaneously. This is achieved using channels, goroutines, etc. Go Language has garbage collection which itself does the memory management and allows the deferred execution of functions.
  * **Go is an open-source programming language focused on simplicity, reliability, and efficiency.**


# Gin-Gonic curd operation without using database

In this code, I am going ot explain about, how we can do curd operation in go lang using gin-gonic framework without using any database. 

We will create a gin-gonic project by Running  this command `go mod init project_name` and `go mod tidy`.

Step 1 :: We will create a struct to define the format of data to be store. 

```go 
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
```

Step 2 :: We will create a array type of variable of `Article` to hold the all data. 

```go
var article []Article
```

Step 3 :: We will create a article and store into article type of varible. 

```go
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
```

step 4 :: We will fetch all the articles

```go
func getAllArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": article,
	})
	return
}
```

step 5 :: We will update a article by article id. 

```go 
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
```

step 6 :: We will create a function, which will delete the existing article or update the existing article

```go
func _deleteArticleByid(id string) {
	for index, art := range article {
		if art.ID == id {
			article = append(article[:index], article[index+1:]...)
			break
		}
	}
}
```

step 7 :: We will delete an article by using article id 

```go
func deleteArticle(c *gin.Context) {
	id := c.Param("id")
	_deleteArticleByid(id)
	c.JSON(http.StatusOK, gin.H{
		"data": article,
	})
	return
}
```

Step 9 :: We will fetch the article by using article id

```go
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

```

Step 10 :: We will handle those rwquest using gin-gonic 

```go
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
```

Step 11 ::: At he end, We will run our project 

```go
func main() {
	handleRequest()
}

```
