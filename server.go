package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func saveMessageView(c *gin.Context) {
	message := c.PostForm("message")
	key := keyBuilder.Get()
	keeper.Set(key, message)
	c.HTML(http.StatusOK, "key.html", gin.H{"key": fmt.Sprintf("http://%s/%s", c.Request.Host, key)})
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html", "templates/key.html")
	router.GET("/", indexView)
	router.POST("/", saveMessageView)
	return router
}
func main() {
	router := getRouter()
	router.Run("localhost:8080")
}

// GO RUN .
// TODO TEST 1 57 07
