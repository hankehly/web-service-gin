package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbum)
	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, albums)
}

// curl -i -X POST -H 'Content-Type: application/json' \
//   -d '{"id":"4","title":"test-title","artist":"test-artist","price":99.99}'  http://localhost:8080/albums
func postAlbums(c *gin.Context) {
	var newAlbum album
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, newAlbum)
}

func getAlbum(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}

	message := fmt.Sprintf("Album where id=%s not found", id)
	c.JSON(http.StatusNotFound, gin.H{"message": message})
}
