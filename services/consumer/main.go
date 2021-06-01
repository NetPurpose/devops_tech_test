package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

const holdOff = 2 * time.Second

func getStream(s *Session) <-chan amqp.Delivery {
	for {
		stream, err := s.Stream(os.Getenv("INBOUND"))
		if err != nil {
			<-time.After(holdOff)
			continue
		}
		return stream
	}
}

func listen(s *Session) {
	log.Println("Getting stream")
	stream := getStream(s)
	for d := range stream {
		reply := fmt.Sprintf("Thanks for that (%s) - All your base are belong to us", string(d.Body))
		log.Printf("Replying to message '%s'", reply)
		s.Push([]byte(reply))
		d.Ack(false)
		log.Println("Done!")
	}

}

func main() {
	forever := make(chan bool)
	s := New(os.Getenv("OUTBOUND"))
	go listen(s)
	<-forever
}
