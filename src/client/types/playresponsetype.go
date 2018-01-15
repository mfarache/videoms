package types

import (
	"fmt"
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

func TracePlayResponse(record PlayResponse) {
	fmt.Println("Genre     = ", record.Genre)
	fmt.Println("ID		   = ", record.ID)
	fmt.Println("Summary   = ", record.Summary)
	fmt.Println("Title	   = ", record.Title)
	fmt.Println("UserViews     = ", record.UserViews)
	fmt.Println("AssetViews	   = ", record.AssetViews)
	fmt.Println("URL     = ", record.Url)
}
