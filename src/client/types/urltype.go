package types

import (
	"fmt"
)

type Url struct {
	ID  string `json:"id"`
	Url string `json:"url"`
}

func TraceURL(record Url) {
	fmt.Println("ID		   = ", record.ID)
	fmt.Println("URL     = ", record.Url)
}
