package types

import (
	"fmt"
)

type Stats struct {
	UserViews  int `json:"userviews"`
	AssetViews int `json:"assetviews"`
}

func TraceStats(recordStats Stats) {
	fmt.Println("UserViews     = ", recordStats.UserViews)
	fmt.Println("AssetViews	   = ", recordStats.AssetViews)
}
