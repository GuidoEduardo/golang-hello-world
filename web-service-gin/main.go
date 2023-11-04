package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)


type Album struct {
	ID			uint64
	ExternalID	string
	Title		string
	Artist		string
	Price		float64
}

var albums = []Album{
	{ID: 1, ExternalID: uuid.NewString(), Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, ExternalID: uuid.NewString(), Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, ExternalID: uuid.NewString(), Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/album/:id", getAlbum)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8000")
}	

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(ctx *gin.Context) {
	var newAlbum Album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbum(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64); if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "album id must be a uint"})
		return
	}

	for _, album := range albums {
		if album.ID == id {
			ctx.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}