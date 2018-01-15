package queue

import (
	"log"

	"github.com/nats-io/go-nats"
)

func Connect(server string) (*nats.Conn, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Connection to NATS done")
	}
	return nc, err
}
