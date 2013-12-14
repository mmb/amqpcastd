package amqpcast

import (
	"code.google.com/p/go.net/websocket"
)

type connection struct {
	ws       *websocket.Conn
	outbound chan string
}

func (c *connection) write() {
	for message := range c.outbound {
		err := websocket.Message.Send(c.ws, message)
		if err != nil {
			break
		}
	}
}

func (c *connection) read() {
	for {
		var message string
		err := websocket.Message.Receive(c.ws, &message)
		if err != nil {
			break
		}
	}
}

func (c *connection) close() {
	c.ws.Close()
	close(c.outbound)
}
