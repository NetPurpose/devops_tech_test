package main

import "github.com/gin-gonic/gin"
import "log"

type Message struct {
	Body string `form:"body" json:"body" binding:"required"`
}

func ping(c *gin.Context, s *Session) {
	var msg string
	if s.IsReady {
		msg = "pong - channel is up!"
	} else {
		msg = "pong - channel is down :("
	}

	c.JSON(200, gin.H{"message": msg})
}

func send(c *gin.Context, s *Session) {
	var json Message
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
	}
	log.Printf("Attempting to Send message %s", string(json.Body))
	err := s.Push([]byte(json.Body))
	if err != nil {
		c.JSON(422, gin.H{"error": err.Error()})
	}
	c.JSON(200, gin.H{"message": "sent"})
}

func receive(c *gin.Context, s *Session) {
	msg, err := s.GetOne("Out")
	if err != nil {
		c.JSON(404, gin.H{"message": "Queue not ready yet"})
		return
	}
	if msg == nil {
		c.JSON(404, gin.H{"message": "No messages in queue"})
		return
	}

	c.JSON(200, gin.H{"message": string(msg)})

}

func main() {
	s := New("dummy")
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		ping(c, s)
	})
	r.POST("/send", func(c *gin.Context) {
		send(c, s)
	})
	r.GET("/receive", func(c *gin.Context) {
		receive(c, s)
	})
	r.Run()
}
