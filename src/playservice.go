package main

import (
	"client"
	"client/types"
	"flag"
	"fmt"
	"log"
	"queue"

	"github.com/gin-gonic/gin"
)

type PlayResponse struct {
	Genre      string `json:"genre"`
	ID         string `json:"id"`
	Summary    string `json:"summary"`
	Title      string `json:"title"`
	UserViews  int    `json:"userviews"`
	AssetViews int    `json:"assetviews"`
	Url        string `json:"url"`
}

var (
	port = flag.String("port", "8080", "Listeninng port")
)

func init() {
	flag.Parse()
}

func createPlayRouter() *gin.Engine {
	r := gin.Default()

	r.POST("play/asset/:assetId/user/:userId", handlePlayContent)
	return r
}

func handlePlayContent(c *gin.Context) {
	assetId := c.Params.ByName("assetId")
	userId := c.Params.ByName("userId")

	recordContent := client.GetAssetDetail(assetId)
	types.TraceContent(recordContent)

	recordUrl := client.GetAssetURL(assetId)
	types.TraceURL(recordUrl)

	var nc, err = queue.Connect("server")
	if err != nil {
		log.Fatal(err)
	} else {
		message := fmt.Sprintf("%s:%s", assetId, userId)
		queue.PublishMessage(nc, "events", message)

	}

	recordStats := client.GetStats(assetId, userId)
	types.TraceStats(recordStats)

	c.JSON(200, gin.H{"Id": recordContent.ID,
		"title":      recordContent.Title,
		"summary":    recordContent.Summary,
		"assetviews": recordStats.UserViews,
		"userviews":  recordStats.AssetViews,
		"url":        recordUrl.Url})
}

func main() {
	serverPort := fmt.Sprintf(":%s", *port)
	r := createPlayRouter()
	r.Run(serverPort)
}
