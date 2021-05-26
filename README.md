# rabbitmq-publisher-example

## How to test

Launch docker with RabbitMQ

```
docker-composer up -d
```

After that build the example and run

```
go build cmd/publish_example.go

./publish_example
```

Remember rabbitmq admin url http://localhost:15672/ (guest:guest)