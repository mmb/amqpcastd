Listen to AMQP messages and send them to websockets clients.

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
