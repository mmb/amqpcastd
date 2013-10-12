package main

import (
	"flag"
	"log"

	"github.com/mmb/amqpcast"
)

func main() {
	var amqpUrl = flag.String("amqp-url", "amqp://localhost:5672/",
		"AMQP server URL")
	var amqpExchange = flag.String("amqp-exchange", "test", "AMQP exchange")
	var amqpKey = flag.String("amqp-key", "test", "AMQP routing key")
	var httpListen = flag.String("http-listen", ":12345", "HTTP listen host and port")

	flag.Parse()

	var cstr = amqpcast.Caster{
		Connections: make(map[*amqpcast.Connection]bool),
		Create:      make(chan *amqpcast.Connection),
		Destroy:     make(chan *amqpcast.Connection),
		Outbound:    make(chan string, 256),
	}

	amqpcast.InitHttp(httpListen, &cstr)
	amqpcast.InitAmqp(amqpUrl, amqpExchange, amqpKey, &cstr)

	for {
		select {
		case c := <-cstr.Create:
			log.Printf("new client")
			cstr.Connections[c] = true
		case c := <-cstr.Destroy:
			log.Printf("client closed")
			delete(cstr.Connections, c)
			c.Ws.Close()
			close(c.Outbound)
		case m := <-cstr.Outbound:
			for c, _ := range cstr.Connections {
				c.Outbound <- m
			}
		}
	}
}
