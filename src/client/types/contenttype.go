package types

import (
	"fmt"
)

type Content struct {
	Genre   string `json:"genre"`
	ID      string `json:"id"`
	Summary string `json:"summary"`
	Title   string `json:"title"`
}

func TraceContent(recordContent Content) {
	fmt.Println("Genre     = ", recordContent.Genre)
	fmt.Println("ID		   = ", recordContent.ID)
	fmt.Println("Summary   = ", recordContent.Summary)
	fmt.Println("Title	   = ", recordContent.Title)
}
