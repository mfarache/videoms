package queue

import (
	"log"

	"github.com/nats-io/go-nats"
)

func PublishMessage(nc *nats.Conn, subject string, message string) {

	if nc != nil {
		defer nc.Close()

		msg := []byte(message)

		nc.Publish(subject, msg)
		nc.Flush()

		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		} else {
			log.Printf("Published [%s] : '%s'\n", subject, msg)
		}
	} else {
		log.Printf("The NC connection is not ok\n")
	}

}
