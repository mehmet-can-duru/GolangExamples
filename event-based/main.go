package main

import (
	"log"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS server
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Fatal(err)
	}

	// "user-created" trigger an event called
	err = nc.Publish("user-created", []byte("New user created"))
	if err != nil {
		log.Fatal(err)
	}

	// close connection
	nc.Close()
}
