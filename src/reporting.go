package main

import (
	"flag"
	"fmt"

	"github.com/gin-gonic/gin"
)

type mapKey struct {
	Key string
}

var (
	port         = flag.String("port", "8082", "Listeninng port")
	userviewsDB  = make(map[mapKey]int)
	assetviewsDB = make(map[mapKey]int)
)

func init() {
	flag.Parse()
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.POST("events/asset/:assetId/user/:userId", registerViewPerUserAndAsset)

	r.GET("events/user/:userId", viewsPerUser)

	r.GET("events/asset/:assetId", viewsPerAsset)

	return r
}

func registerViewPerUserAndAsset(c *gin.Context) {
	assetId := c.Params.ByName("assetId")
	assetCounter, ok := assetviewsDB[mapKey{assetId}]
	if ok {
		assetviewsDB[mapKey{assetId}] = assetCounter + 1
	} else {
		assetviewsDB[mapKey{assetId}] = 1
	}

	userId := c.Params.ByName("userId")
	userCounter, ok := userviewsDB[mapKey{userId}]
	if ok {
		userviewsDB[mapKey{userId}] = userCounter + 1
	} else {
		userviewsDB[mapKey{userId}] = 1
	}
	c.JSON(200, gin.H{"userviews": userviewsDB[mapKey{userId}],
		"assetviews": assetviewsDB[mapKey{assetId}]})
}

func viewsPerUser(c *gin.Context) {
	userId := c.Params.ByName("userId")

	value, ok := userviewsDB[mapKey{userId}]
	if ok {
		c.JSON(200, gin.H{"id": userId,
			"views": value})
	} else {
		c.JSON(404, gin.H{"id": userId, "status": "User not found"})
	}
}

func viewsPerAsset(c *gin.Context) {
	assetId := c.Params.ByName("assetId")
	value, ok := assetviewsDB[mapKey{assetId}]
	if ok {
		c.JSON(200, gin.H{"id": assetId,
			"views": value})
	} else {
		c.JSON(404, gin.H{"id": assetId, "status": "Asset not found"})
	}
}

func main() {
	serverPort := fmt.Sprintf(":%s", *port)
	r := setupRouter()
	r.Run(serverPort)
}
