package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// define a struct to enter data
type album struct {
	ID       string `json:"id"`
	title    string `json:"title"`
	artist   string `json:"artist"`
	numSongs int    `json:"price"`
}

var albums = []album{
	{ID: "1", title: "In a Minute", artist: "Lil Baby", numSongs: 1},
	{ID: "2", title: "Scropion", artist: "Drake", numSongs: 25},
	{ID: "3", title: "Blame it on baby", artist: "DaBaby", numSongs: 24},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
