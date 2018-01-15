package client

import (
	"client/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const HOST_PORT_CONTENT_SERVICE = "localhost:8081"

func GetAssetDetail(assetId string) types.Content {
	var record types.Content
	//Get content
	url := fmt.Sprintf("http://%s/content/asset/%s", HOST_PORT_CONTENT_SERVICE, assetId)
	// Build the request

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Obtaining content failed: ", err)
		return record
	}

	client := &http.Client{}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return record
	}

	defer response.Body.Close()

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record
}
