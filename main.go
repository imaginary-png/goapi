package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// responds with list of albums, JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "This is not a real album", Artist: "Fake Name", Price: 99.83},
	{ID: "5", Title: "Floating through the Clouds", Artist: "Starry Fuzz", Price: 44.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
