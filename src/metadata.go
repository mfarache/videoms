package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

type typeContent struct {
	title   string
	genre   string
	summary string
}

type keyContent struct {
	Key string
}

var (
	port    = flag.String("port", "8081", "Listeninng port")
	assetDB = make(map[keyContent]typeContent)
)

func init() {
	flag.Parse()
}

func createAssetDB() {
	assetDB[keyContent{"123"}] = typeContent{"It", "horror", "A very scary movie"}
	assetDB[keyContent{"456"}] = typeContent{"Terminator", "action", "Summary for terminator"}
	fmt.Printf("%+v\n", assetDB)
}

func createVideoContentRouter() *gin.Engine {

	createAssetDB()

	r := gin.Default()

	r.GET("content/asset/:assetId", handleGetContent)
	r.POST("content/asset/:assetId", handlePostContent)

	return r
}

func handleGetContent(c *gin.Context) {

	id := c.Params.ByName("assetId")

	value, ok := assetDB[keyContent{id}]
	if ok {
		c.JSON(200, gin.H{"id": id,
			"title":   value.title,
			"genre":   value.genre,
			"summary": value.summary})
	} else {
		c.JSON(404, gin.H{"id": id, "status": "Asset not found"})
	}
}

func handlePostContent(c *gin.Context) {

	var json typeContent
	id := c.Params.ByName("assetId")
	if c.Bind(&json) == nil {
		assetDB[keyContent{id}] = json
		c.JSON(200, gin.H{"status": "ok"})
	}

}

func main() {
	serverPort := fmt.Sprintf(":%s", *port)
	r := createVideoContentRouter()
	r.Run(serverPort)

}
