package main

import "time"
import "github.com/streadway/amqp"
import "fmt"

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
	stream := getStream(s)
	for d := range stream {
		d.Ack(false)
		reply := fmt.Sprintf("Thanks for that (%s) - All your base are belong to us", string(d.Body))
		s.Push([]byte(reply))
	}

}

func main() {
	forever := make(chan bool)
	s := New("dummy")
	go listen(s)
	<-forever
}
