package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Album struct {
	ID     string  `json:"id"`
	TITLE  string  `json:"title"`
	ARTIST string  `json:"artist"`
	PRICE  float64 `json:"price"`
}

var albums = []Album{
	{"1", "mona lisa", "lil wayne", 25.99},
	{"2", "holy", "justin beiber", 59.00},
	{"3", "mbona", "khaligraph", 20.65},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

// get albums responds with list of all albums as json
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum Album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
