package amqpcast

import (
	"log"
	"sync"
)

type caster struct {
	sync.RWMutex
	connections map[*connection]bool
	create      chan *connection
	destroy     chan *connection
	outbound    chan string
}

func NewCaster() *caster {
	return &caster{
		connections: make(map[*connection]bool),
		create:      make(chan *connection),
		destroy:     make(chan *connection),
		outbound:    make(chan string, 256),
	}
}

func (cstr *caster) Run() {
	for {
		select {
		case c := <-cstr.create:
			log.Printf("new client")
			cstr.createConnection(c)
		case c := <-cstr.destroy:
			log.Printf("client closed")
			cstr.destroyConnection(c)
		case m := <-cstr.outbound:
			cstr.distributeMessage(m)
		}
	}
}

func (cstr *caster) createConnection(c *connection) {
	cstr.Lock()
	cstr.connections[c] = true
	cstr.Unlock()
}

func (cstr *caster) destroyConnection(c *connection) {
	cstr.Lock()
	delete(cstr.connections, c)
	cstr.Unlock()
	c.close()
}

func (cstr *caster) distributeMessage(m string) {
	cstr.RLock()
	for c, _ := range cstr.connections {
		c.outbound <- m
	}
	cstr.RUnlock()
}
