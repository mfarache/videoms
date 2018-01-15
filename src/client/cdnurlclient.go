package client

import (
	"client/types"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const HOST_PORT_URL_SERVICE = "localhost:8083"

func GetAssetURL(assetId string) types.Url {
	var record types.Url
	url := fmt.Sprintf("http://%s/url/asset/%s", HOST_PORT_URL_SERVICE, assetId)
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
	if err := json.NewDecoder(response.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	return record
}
