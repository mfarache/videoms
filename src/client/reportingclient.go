package client

import (
	"client/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const HOST_PORT_REPORTING_SERVICE = "localhost:8082"

func GetStats(assetId string, userId string) types.Stats {

	var record types.Stats
	url := fmt.Sprintf("http://%s/events/asset/%s/user/%s", HOST_PORT_REPORTING_SERVICE, assetId, userId)

	// Build the request
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		log.Fatal("Reporting views failed: ", err)
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
