# RabbitMQ

Imagine You (the publisher) send a marriage invitation (message) to the post office (exchange). The post office (exchange) looks at the Address (routing key) and decides which Flat's PostBox (queue) should get the invitation. The Flat's PostBox (queue) keeps the invitation until someone (the consumer) comes to read it.

The fundamental goal is to decouple producers and consumers, enabling asynchronous and reliable message processing. In RabbitMQ, an exchange is a routing mechanism that receives messages from producers (publishers) and routes them to one or more queues based on certain rules or bindings.  The publisher sends messages to an exchange (not directly to a queue).

In RabbitMQ, queues should be declared in the consumer side. The publisher does not need to declare queues, but instead, it publishes messages to an exchange. The exchange is responsible for routing the messages to the appropriate queues based on the routing key and binding configuration.

Publisher will create Connection then from Connection it will create Channel then from channel it call ExchangeDeclare then it sends messages using Publish function of channel.

Consumer will create Connection then from Connection it will create Channel then from channel it call ExchangeDeclare then it call QueueDeclare then it Binds the Queue and Exchange by QueueBind function of Channel. Then call Consume function of channel.

In Topic Exchange Publisher sends messages to specific queues based on routing pattern specified by the routing key. The routing key can include wildcards (* for a single word and # for multiple words).

Example :- 

Publisher :
```go
func main() {
	// Establish a connection to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect: %s", err)
	}
	defer conn.Close() // Ensure connection is closed when the function exits
	// Create a channel to communicate with RabbitMQ
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel: %s", err)
	}
	defer ch.Close() // Ensure channel is closed when the function exits
	// Declare a topic exchange named "topic_logs"
	err = ch.ExchangeDeclare(
		"topic_logs",    // Exchange name
		"topic",     // Exchange type (topic exchange)
		true,        // Durable: The exchange will survive server restarts
		false,       // Auto-deleted: The exchange won't be deleted when no consumers are connected
		false,       // Internal: The exchange is not internal to RabbitMQ
		false,       // No-wait: Don't wait for confirmation when declaring the exchange
		nil,         // Additional arguments (none in this case)
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %s", err)
	}
	// Publish some messages
	// Routing key that matches the pattern for consumers (in this case, it will match "animal.cat")
	messages := []struct {
		routingKey string
		body       string
	}{
		{"animal.cat", "Cat message"},
		{"animal.dog", "Dog message"},
		{"animal.bird", "Bird message"},
	}
	for _, msg := range messages {
		err = ch.Publish(
			exchange, msg.routingKey, false, false,
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(msg.body),
			},
		)
		if err != nil {
			log.Fatalf("Failed to publish message: %s", err)
		}
	}
	fmt.Println("Message sent:", messages)
}
```
Consumer :
```go
// Consumer helper-function to handle message processing
func consumeMessages(ch *amqp.Channel, wg *sync.WaitGroup, queueName string, prefetchCount int) {
	defer wg.Done()
	// Set the prefetch count: This ensures the consumer will not receive more than `prefetchCount` messages at once
	err := ch.Qos(prefetchCount, 0, false) // The consumer will only receive 5 messages at once
	if err != nil {
		log.Fatalf("Failed to set Qos: %s", err)
	}
	// Start consuming messages from the queue
	msgs, err := ch.Consume(
		queueName, // Queue name
		"",        // Consumer tag (empty string means RabbitMQ will generate a tag)
		false,     // Auto-acknowledge messages (false means manual ack is required)
		false,     // Exclusive: This consumer is not exclusive to the connection
		false,     // No-local: Don't deliver messages published by the same connection
		false,     // No-wait: Don't wait for acknowledgment
		nil,       // Additional arguments (none in this case)
	)
	if err != nil {
		log.Fatalf("Failed to consume messages: %s", err)
	}
	// Process each message
	for msg := range msgs {
		// Simulate processing of the message (we can replace this with actual business logic)
		fmt.Printf("Processing message: %s\n", msg.Body)
		time.Sleep(1 * time.Second) // Simulate processing time (can be replaced with actual logic)
		// Acknowledge the message to let RabbitMQ know it's been processed
		msg.Ack(false)
	}
}
```
We are using topic exchange (topic_logs) with messages being published to it with routing keys such as animal.cat, animal.dog, animal.bird. The consumers are bound to receive messages that match the animal.* pattern, so they will get messages with keys like animal.cat, animal.dog, and so on.
```go
// CONSUMER
func main() {
	// Establish a connection to RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect: %s", err)
	}
	defer conn.Close() // Ensure connection is closed when the function exits
	// Create a channel to communicate with RabbitMQ
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to create channel: %s", err)
	}
	defer ch.Close() // Ensure channel is closed when the function exits
	// Declare a topic exchange named "topic_logs"
	exchange := "topic_logs"
	err = ch.ExchangeDeclare(
		exchange,    // Exchange name
		"topic",     // Exchange type (topic exchange)
		true,        // Durable: The exchange will survive server restarts
		false,       // Auto-deleted: The exchange won't be deleted when no consumers are connected
		false,       // Internal: The exchange is not internal to RabbitMQ
		false,       // No-wait: Don't wait for confirmation when declaring the exchange
		nil,         // Additional arguments (none in this case)
	)
	if err != nil {
		log.Fatalf("Failed to declare exchange: %s", err)
	}
	// Declare a queue (anonymous, will be deleted after the consumer disconnects)
	queue, err := ch.QueueDeclare(
		"",    // Empty queue name (RabbitMQ generates a random queue name)
		false, // Durable: The queue won't survive server restarts
		true,  // Delete when unused: The queue will be deleted when no consumers are connected
		true,  // Exclusive: The queue is exclusive to this connection
		false, // No-wait: Don't wait for confirmation when declaring the queue
		nil,   // Additional arguments (none in this case). will use for dead letter exchange
	)
	if err != nil {
		log.Fatalf("Failed to declare queue: %s", err)
	}
	// Set prefetch count to 5 (each consumer will only receive 5 messages at a time)
	prefetchCount := 5
	// Bind the queue to the exchange with a routing key pattern "animal.*"
	routingKey := "animal.*"
	err = ch.QueueBind(
		queue.Name,   // Queue name
		routingKey,   // Routing key pattern
		exchange,     // Exchange to bind to
		false,        // No-wait: Don't wait for confirmation when binding the queue
		nil,          // Additional arguments (none in this case)
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %s", err)
	}

	// Use a wait group to wait for all consumers to finish processing
	var wg sync.WaitGroup
	numConsumers := 3 // Number of consumers to start concurrently
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go consumeMessages(ch, &wg, queue.Name, prefetchCount)
	}

	// Wait for all consumers to finish processing
	wg.Wait()
}
```
In Fanout Exchange Publisher sends messages to all Queues connected to Fanout Exchanges. 
```go
// Declare fanout exchange
ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)

// Publish a message (routing key is ignored for fanout)
err = ch.Publish("logs", "", false, false, amqp.Publishing{
    ContentType: "text/plain",
    Body:        []byte("Fanout exchange message"),
})
```
```go
// Declare fanout exchange
ch.ExchangeDeclare("logs", "fanout", true, false, false, false, nil)

// Declare an anonymous queue
queue, err := ch.QueueDeclare("", false, true, true, false, nil)
if err != nil {
    log.Fatalf("Failed to declare queue: %s", err)
}

// Bind the queue to the exchange (routing key is ignored for fanout)
err = ch.QueueBind(queue.Name, "", "logs", false, nil)
```
In Direct Exchange Publisher sends message to specific Queues based on a Exact match of Routing key.
```go
// Declare direct exchange
ch.ExchangeDeclare("direct_logs", "direct", true, false, false, false, nil)

// Publish a message with a routing key (e.g., "error")
err = ch.Publish("direct_logs", "error", false, false, amqp.Publishing{
    ContentType: "text/plain",
    Body:        []byte("Error in processing the task"),
})
```
```go
// Declare direct exchange
ch.ExchangeDeclare("direct_logs", "direct", true, false, false, false, nil)

// Declare an anonymous queue
queue, err := ch.QueueDeclare("", false, true, true, false, nil)
if err != nil {
    log.Fatalf("Failed to declare queue: %s", err)
}

// Bind the queue to the exchange with the routing key "error"
err = ch.QueueBind(queue.Name, "error", "direct_logs", false, nil)
```

In Header Exchnage Publisher sends messages to all queues based on Message Headers rather than Routing Key.
```go
// Declare header exchange
ch.ExchangeDeclare("header_logs", "headers", true, false, false, false, nil)

// Define message headers
headers := amqp.Table{
    "X-Color": "Red",  // Set a header key-value pair
}

// Publish a message with headers
err = ch.Publish("header_logs", "", false, false, amqp.Publishing{
    Headers:     headers,
    ContentType: "text/plain",
    Body:        []byte("Header exchange message with color Red"),
})
```
```go
// Declare header exchange
ch.ExchangeDeclare("header_logs", "headers", true, false, false, false, nil)

// Declare an anonymous queue
queue, err := ch.QueueDeclare("", false, true, true, false, nil)
if err != nil {
    log.Fatalf("Failed to declare queue: %s", err)
}

// Bind the queue using header matching (e.g., "X-Color" equals "Red")
headers := amqp.Table{
    "X-Color": "Red",  // Match this header for routing
}
err = ch.QueueBind(queue.Name, "", "header_logs", false, headers)

```
- **Durable Queue**: Survives server restarts, messages are not lost. In `ch.QueueDeclare()` function we have one boolean argument `durable` by which we can set it.
- **Transient Queue**: Does not survive a server restart.  In `ch.QueueDeclare()` function we have one boolean argument `durable` by which we can unset it and make it Transient.
- **Exclusive Queue**: Only used by one consumer and deleted when that consumer disconnects. In `ch.QueueDeclare()` function we have one boolean argument `exclusive` by which we can set it. 
- **Auto-Delete Queue**: Deletes itself when no consumers are using it. In `ch.QueueDeclare()` function we have one boolean argument `auto-delete` by which we can set it. 

- **Dead Letter Queues**: Special queues for storing messages that can't be delivered (e.g., after too many retries).
```go
// Declare a dead letter exchange
	args := amqp.Table{
		"x-dead-letter-exchange": "dlx_exchange",
	}
	_, err = ch.QueueDeclare(
		"normal_queue",
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		args,  // arguments for DLX
	)
```
- **Automatic Acknowledgement**: The message is automatically confirmed once sent to the consumer. This is the default behavior. In `ch.Consume()` function we have one boolean argument `auto-ack` by which we can set it.

- **Manual Acknowledgement**: The consumer confirms it has received and processed a message. In `ch.Consume()` function we have one boolean argument `auto-ack` by which we can unset it. In this case, the message is not considered acknowledged until we explicitly call the Ack method. `msg.Ack(false)`.

- **Negative Acknowledgement (Nack)**: When a message is rejected by a consumer, RabbitMQ can retry or discard it. In this case, we can use Nack to reject the message, and optionally, we can requeue it for another attempt. `msg.Nack(false, true)` true to requeue the message

- **Clustering**: Running RabbitMQ on multiple servers to spread the load and ensure it's always available.
- **Mirrored Queues**: Copying queues to other servers in the cluster so that if one server fails, messages are still available.
- **Federation**: Allowing RabbitMQ instances in different locations to share messages.
- **Prefetch Count**: Limit how many messages a consumer can process at once to avoid overwhelming it.
- **Concurrency**: Running many consumers in parallel to handle more messages.
- **Flow Control**: Managing how messages are processed to avoid overloading RabbitMQ.
- **Memory and Disk Usage**: Keep track of resource usage to prevent performance issues.


TLS encrypts the communication between clients and RabbitMQ servers, preventing eavesdropping or interception by third parties. Without TLS, messages and credentials sent over the network can be easily captured in plaintext, especially in a shared or insecure network environment (e.g., public Wi-Fi or unsecured VPNs).

TLS helps authenticate the server. When a client connects to RabbitMQ over a TLS-secured connection, it verifies the RabbitMQ server's certificate (using a CA certificate), ensuring the client is connecting to the right server and not an imposter.

We can configure RabbitMQ to request client certificates as part of the TLS handshake. This allows for mutual authentication, meaning both the server and the client verify each other's identity.

Ensure our RabbitMQ server is set up to use TLS. You can modify the RabbitMQ configuration file (rabbitmq.conf) to enable TLS.
```bash
listeners.tls.default = 5671
ssl_options.cacertfile = /path/to/ca_certificate.pem
ssl_options.certfile = /path/to/server_certificate.pem
ssl_options.keyfile = /path/to/server_key.pem
ssl_options.verify = verify_peer
ssl_options.fail_if_no_peer_cert = true
```
listeners.tls.default: The default port for secure connections (5671 is the default for TLS).

ssl_options.cacertfile: The certificate authority's (CA) certificate.

ssl_options.certfile: The RabbitMQ server's certificate.

ssl_options.keyfile: The private key associated with the RabbitMQ server certificate.

ssl_options.verify: Controls whether to verify client certificates.

ssl_options.fail_if_no_peer_cert: If true, RabbitMQ will reject connections without a valid certificate.


```go
	// Load the CA cert
	caCert, err := os.ReadFile("/path/to/ca_certificate.pem")
	if err != nil {
		log.Fatalf("Failed to read CA cert: %v", err)
	}

	// Create a certificate pool from the CA cert
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(caCert)

	// Create the TLS config
	tlsConfig := &tls.Config{
		RootCAs: certPool,
	}

	// RabbitMQ URL
	rabbitmqURL := "amqps://user:password@our.rabbitmq.server:5671/"

	// Dial the connection with TLS
	conn, err := amqp.DialTLS(rabbitmqURL, tlsConfig)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()
```

Client Certificates: If RabbitMQ requires a client certificate, we can load the certificate and private key into the tls.Config object.
```go
clientCert, err := tls.LoadX509KeyPair("/path/to/client_cert.pem", "/path/to/client_key.pem")
if err != nil {
    log.Fatalf("Failed to load client cert: %v", err)
}

tlsConfig := &tls.Config{
    RootCAs:      certPool,
    Certificates: []tls.Certificate{clientCert},
}
```