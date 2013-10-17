This server listens to a routing key on an AMQP exchange and broadcasts all
messages received to currently connected WebSocket clients.

It can be used as plumbing to shove data into a queue and get it into users'
browsers.

The server starts an example page but you could create a WebSocket connection
to the server from your own page.

Assuming you have RabbitMQ installed and listening on localhost:5672:

```sh
go install

rabbitmqadmin declare exchange name=test type=direct

amqpcastd \
--amqp-url amqp://localhost:5672/ \
--amqp-exchange=test \
--amqp-key=test

open http://localhost:12345/

rabbitmqadmin publish exchange=test routing_key=test payload="hello world"
```

Thanks to @garyburd for this helpful gist:

https://gist.github.com/garyburd/1316852
