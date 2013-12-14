package main

import (
	"flag"

	"github.com/mmb/amqpcastd/amqpcast"
)

func main() {
	var amqpUrl = flag.String("amqp-url", "amqp://localhost:5672/",
		"AMQP server URL")
	var amqpExchange = flag.String("amqp-exchange", "test", "AMQP exchange")
	var amqpKey = flag.String("amqp-key", "test", "AMQP routing key")
	var httpListen = flag.String("http-listen", ":12345", "HTTP listen host and port")

	flag.Parse()

	cstr := amqpcast.NewCaster()

	amqpcast.InitHttp(httpListen, cstr)
	amqpcast.InitAmqp(amqpUrl, amqpExchange, amqpKey, cstr)

	cstr.Run()
}
