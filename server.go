package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func getRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html")
	router.GET("/", indexView)
	return router
}
func main() {
	router := getRouter()
	router.Run("localhost:8080")
}