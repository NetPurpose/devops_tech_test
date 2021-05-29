package main

import (
	"errors"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

type callback = func() error

type Session struct {
	name    string
	conn    *amqp.Connection
	channel *amqp.Channel
	done    chan bool
	IsReady bool
}

const (
	retryHoldoff = 5 * time.Second
)

func New(name string) *Session {
	addr := os.Getenv("AMQP_URL")
	if addr == "" {
		log.Fatalf("No address provided")
	}
	s := Session{
		name: name,
		done: make(chan bool),
	}
	go s.init(addr)
	return &s
}

func (s *Session) init(addr string) {
	if err := s.connect(addr); err != nil {
		log.Fatalf("Failed to connect to rabbitmq")
	}
	if err := s.setupChannel(); err != nil {
		log.Fatalf("Failed to establish channel")
	}
	s.IsReady = true
}

func (s *Session) connect(addr string) error {
	log.Println("Connecting to rabbitmq")
	return tryUntilComplete(func() error {
		ch, err := amqp.Dial(addr)
		if err != nil {
			return err
		}
		s.conn = ch
		return nil
	})
}

func (s *Session) setupChannel() error {
	log.Println("Establishing channel")
	return tryUntilComplete(func() error {
		ch, err := s.conn.Channel()
		if err != nil {
			return err
		}
		s.channel = ch
		_, err = ch.QueueDeclare(
			s.name,
			false,
			false,
			false,
			false,
			nil)
		if err != nil {
			return err
		}
		return nil
	})
}

func (s *Session) Push(msg []byte) error {
	log.Println("Sending message")
	if !s.IsReady {
		return errors.New("connection not ready")
	}
	return s.channel.Publish(
		"",
		s.name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		})

}

func (s *Session) GetOne() ([]byte, error) {
	d, ok, err := s.channel.Get(s.name, false)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, nil
	}
	d.Ack(false)
	return d.Body, nil

}

func (s *Session) Stream(name string) (<-chan amqp.Delivery, error) {
	if !s.IsReady {
		return nil, errors.New("connection not ready")
	}
	return s.channel.Consume(
		name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

func tryUntilComplete(fn callback) error {
	for {
		err := fn()
		if err != nil {
			log.Println("Failed. Retrying...")
			<-time.After(retryHoldoff)
		} else {
			return nil
		}
	}
}
