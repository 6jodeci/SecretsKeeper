package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func writeInternalEror(c *gin.Context) {
	c.HTML(http.StatusInternalServerError, "500.html", gin.H{})
}

func indexView(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func saveMessageView(c *gin.Context, keyBuilder KeyBuilder, keeper Keeper) {
	message := c.PostForm("message")
	key := keyBuilder.Get()
	err := keeper.Set(key, message)
	if err != nil {
		writeInternalEror(c)
		return
	}
	c.HTML(http.StatusOK, "key.html", gin.H{"key": fmt.Sprintf("http://%s/%s", c.Request.Host, key)})
}

func readMessageView(c *gin.Context, keyBuilder KeyBuilder, keeper Keeper) {
	key := c.Param("key")
	msg, err := keeper.Get(key)
	if err != nil {
		if err.Error() == NOT_FOUND_ERROR {
			c.HTML(http.StatusNotFound, "404.html", gin.H{})
			return
		}
		writeInternalEror(c)
		return
	}

	err = keeper.Clean(key)
	if err != nil {
		writeInternalEror(c)
		return
	}
	c.HTML(http.StatusOK, "message.html", gin.H{"message": msg})
}

func buildHandler(fn func(c *gin.Context, keyBuilder KeyBuilder, keeper Keeper), keyBuilder KeyBuilder, keeper Keeper) gin.HandlerFunc {
	return func(c *gin.Context) {
		fn(c, keyBuilder, keeper)
	}
}

func getRouter(keyBuilder KeyBuilder, keeper Keeper) *gin.Engine {
	router := gin.Default()
	router.LoadHTMLFiles("templates/index.html",
		"templates/key.html", "templates/message.html",
		"templates/404.html", "templates/500.html")
	router.GET("/", indexView)
	router.POST("/", buildHandler(saveMessageView, keyBuilder, keeper))
	router.GET("/:key", buildHandler(readMessageView, keyBuilder, keeper))
	return router
}
func main() {
	keyBuilder := getKeyBuilder()
	keeper := getKeeper()
	router := getRouter(keyBuilder, keeper)
	router.Run("localhost:8080")
}

// GO RUN .
// TODO 1 15 35
