package main

import "time"
import "github.com/streadway/amqp"
import "fmt"
import "log"

const holdOff = 2 * time.Second

func getStream(s *Session) <-chan amqp.Delivery {
	for {
		stream, err := s.Stream("dummy")
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
	s := New("Out")
	go listen(s)
	<-forever
}
