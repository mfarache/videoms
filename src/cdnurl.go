package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

type typeURL struct {
	url string
}

type keyURL struct {
	Key string
}

var (
	port  = flag.String("port", "8083", "CDN URL Listeninng port")
	urlDB = make(map[keyURL]typeURL)
)

func init() {
	flag.Parse()
}

func createDB() {
	urlDB[keyURL{"123"}] = typeURL{"path/to/asset123.mp4"}
	fmt.Printf("%+v\n", urlDB)
}

func createCDNUrlContentRouter() *gin.Engine {

	createDB()

	r := gin.Default()

	r.GET("url/asset/:assetId", handleGetURL)
	r.POST("url/asset/:assetId", handlePostURL)

	return r
}

func handleGetURL(c *gin.Context) {

	id := c.Params.ByName("assetId")

	value, ok := urlDB[keyURL{id}]
	if ok {
		c.JSON(200, gin.H{"id": id,
			"url": value.url})
	} else {
		c.JSON(404, gin.H{"id": id, "status": "Asset not found"})
	}
}

func handlePostURL(c *gin.Context) {

	id := c.Params.ByName("assetId")
	var json typeURL
	if c.Bind(&json) == nil {
		urlDB[keyURL{id}] = json
		c.JSON(200, gin.H{"status": "ok"})
	}
}

func main() {
	serverPort := fmt.Sprintf(":%s", *port)
	r := createCDNUrlContentRouter()
	r.Run(serverPort)

}
