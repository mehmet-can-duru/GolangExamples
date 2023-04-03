package main

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}

	// "user-created" listen to an event called
	_, err = nc.Subscribe("user-created", func(msg *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(msg.Data))
	})
	if err != nil {
		log.Fatal(err)
	}

	// stay in endless loop
	select {}
}
