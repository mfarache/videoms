package main

import (
	"log"
	"queue"
)

func main() {

	var nc, err = queue.Connect("server")
	if err != nil {
		log.Fatal(err)
	} else {
		queue.ReceiveMessage(nc, "events")
	}

}
