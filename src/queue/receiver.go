package queue

import (
	"log"
	"runtime"
	"time"

	"github.com/nats-io/go-nats"
)

type typeEvent struct {
	id        string
	eventType string
	when      string
}

type keyEvent struct {
	Key string
}

var (
	eventsDB = make(map[keyEvent]typeEvent)
)

func storeMessage(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, string(m.Data))
	now := time.Now()
	//msg = string(m.Data)
	var msg = "123"
	eventsDB[keyEvent{msg}] = typeEvent{msg, "VIEW", now.Format(time.RFC3339)}
}

func ReceiveMessage(nc *nats.Conn, subject string) {

	if nc != nil {

		i := 0

		nc.Subscribe(subject, func(msg *nats.Msg) {
			i += 1
			storeMessage(msg, i)
		})
		nc.Flush()

		if err := nc.LastError(); err != nil {
			log.Fatal(err)
		}

		log.Printf("Listening on [%s]\n", subject)
	} else {
		log.Printf("The NC connection is not ok\n")
	}

	runtime.Goexit()

}
