# KAFKA

Kafka is an event streaming service (often referred to as a distributed event streaming platform) rather than a typical queue service. Kafka is primarily used for handling massive amounts of data in real-time.

In Kafka, the server is called a `Broker`. A Kafka `Cluster` consists of multiple brokers that work together to handle the data flow efficiently.

Kafka uses `topics` to send and receive messages, which can be thought of as analogous to queues in traditional messaging systems. However, in Kafka, we call them topics. Each topic can have multiple `partitions`. Each partition contains a subset of the messages for a specific topic. Kafka ensures that messages are ordered within a partition, but doesn't guarantee ordering across different partitions in the same topic.

Partitions allow Kafka to scale horizontally. This means that each partition can be processed in parallel by multiple machines, making Kafka capable of handling large-scale data streams.

Kafka retains messages even after they are consumed, and multiple consumers can read the same message at different times.

Multiple consumers can form a `consumer group`. Kafka ensures that only one consumer from the group consumes each message within a topic partition. If there are N partitions for a topic and N consumers in the group, each consumer will be assigned to one partition, and each partition will be read by only one consumer. 

If there are more consumers than partitions, some consumers will be idle because there aren’t enough partitions to assign to them.
If there are more partitions than consumers, Kafka will assign multiple partitions to the same consumer to balance the load. This allows for load balancing and fault tolerance within the consumer group.

Each message in Kafka has an `offset`. The offset is like the page number in a book, it uniquely identifies the position of a message within a partition. Consumers track the offset of the messages they have consumed. If a consumer crashes, it can resume from the last successfully processed offset, ensuring that no messages are missed or read twice.

Kafka follows a `Leader-Follower` model for partition replication. The leader broker is responsible for handling both read and write requests for a partition. Follower brokers replicate data from the leader. If the leader broker crashes, one of the follower brokers is elected to become the new leader, ensuring high availability and fault tolerance. Followers do not handle client requests directly, they simply replicate data from the leader.

In a distributed system, several challenges arise, such as coordination, state consistency, and fault tolerance. Kafka needs to manage its cluster state, keep track of partition leaders, and replicate data across different brokers. This involves keeping all brokers aware of each other’s states, ensuring fault tolerance during broker failures, and managing the distribution of work (messages, partitions, etc.) across the brokers.

Zookeeper is use for managing multiples Kafka brokers in a cluster. Zookeeper is use for :
- Cluster Management and Metadata Storage :  Zookeeper keeps track of which brokers are part of the Kafka cluster.  It stores information about topics, including their configuration, partitioning information, replication, and other metadata.
- Zookeeper plays a critical role in leader election. If the current leader for a partition goes down, Zookeeper is responsible for electing a new leader from the followers.  
-  Zookeeper detects the failure of a broker. Once Zookeeper detects a broker failure, it triggers the election of a new leader for any partitions that were managed by the failed broker
- Kafka often needs to store and manage configuration settings for brokers, topics, and partitions. Zookeeper provides a way to centrally store and manage these configurations in a distributed system.

Without Zookeeper, Kafka would have to implement and maintain this coordination internally, which would add complexity and increase the risk of data loss or duplication in the consumer process.

Kafka retry logic involves attempting to resend a message if it fails to be delivered initially. Kafka's producer and consumer mechanisms support various retry strategies like `retries`,`retry.backoff.ms`, `acks`. 
In Kafka, retries can be done automatically with these settings.

In kafka we have ConfigMap struct where we can initalize 

```go
	// Kafka configuration with retry settings
	config := &kafka.ConfigMap{
		"bootstrap.servers":    "your-kafka-broker:9093", // Broker address
		"acks":                 "all", // Ensure all replicas have acknowledged the message
		"retries":              5,     // Number of retries in case of failure
		"retry.backoff.ms":     500,   // Wait time before retrying (in ms)
		"max.in.flight.requests.per.connection": 1, // Ensure sequential message sending (important for ordering)
		"partitioner":          "murmur2", // Custom partitioning strategy
	}
	// Create a new producer instance
	producer, err := kafka.NewProducer(config)
```

Kafka brokers, producers, and consumers might be deployed in different environments or regions. To ensure that the data transmitted between these entities is secure and not intercepted by unauthorized actors, SSL/TLS is used to encrypt the messages.

Kafka requires authentication mechanisms to ensure that only authorized producers and consumers can interact with the Kafka brokers. This is achieved through SSL/TLS certificates. Without this, anyone could potentially produce or consume messages from your Kafka cluster.


TLS is the successor to SSL and provides stronger encryption and security. TLS is widely used today for securing communication channels over networks, including HTTP (HTTPS), email, and Kafka.

Setting up Kafka with SSL/TLS encryption in a production environment, especially when dealing with microservices, is crucial for ensuring data security and communication integrity. 

- CA Certificate (ca-cert): A trusted certificate authority that verifies the authenticity of the Kafka broker and client certificates.

- Broker Certificate (certificate): A certificate that the Kafka broker uses to prove its identity.

- Private Key (private-key): The private key corresponding to the broker's certificate.

- Client Certificate (certificate): If you're using client authentication (mutual TLS), each client (microservice) will need its own certificate.

Kafka can also authenticate both the client and the server. This is called mutual TLS (mTLS), which adds an extra layer of security by ensuring that both the client (producer/consumer) and the Kafka broker are authenticated using their respective certificates.

The CA certificate is a trusted certificate authority that signs the broker certificate and, optionally, the client certificate. It is used to verify the authenticity of the Kafka broker’s certificate and any client certificates.

Without the CA certificate, a Kafka client would not be able to trust the broker's certificate or verify its authenticity. In a production environment, you typically use certificates signed by a public or private certificate authority (CA). 

Broker certificate allows the Kafka client to verify that it is connecting to the right Kafka broker and not an imposter. If you use mutual authentication, this is essential.  The broker uses this certificate to establish secure connections and ensure encrypted communication. If using mTLS, this certificate is sent during the SSL handshake to prove the identity of the server.

The private key is associated with the broker's certificate. It is used by the Kafka broker to decrypt incoming requests that were encrypted using its public key (from the certificate). The private key is crucial because, during the SSL handshake,

Client is a certificate that each Kafka client (producer or consumer) uses to prove its identity to the Kafka broker during a connection. This is used when mutual TLS (mTLS) is enabled.

Using mTLS, Kafka can restrict access to certain producers and consumers, allowing only authenticated clients to publish and consume messages.


## Steps

- Create Privete key by OPENSSL command choose algoriyh as RSA.
- Generate Certificates. Need to Sign those certificates by CA.
- Kafka Broker need to onfigured to support SSL/TLS. For this we can modify the server.properties file.
- When we call kafka.NewProducer(&kafka.ConfigMap{...} and kafka.NewConsumer(&kafka.ConfigMap{..}  we need to pass security protocol ssl/tls, CA certificate location, Client certificate location, Client private key location, client pasword etc.

In a microservices architecture, each service that communicates with Kafka will need to use the above configuration , Service 1 Needs access to ca.crt, kafka-client.crt, and kafka-client.key. Similarly, Service 2 will need the same certificates (or its own unique client certificates).

Microservices can access the necessary certificates (you may want to store them securely, e.g., in a vault, or use Kubernetes secrets in a containerized environment).

A Key Management System (KMS) is crucial for managing certificates, private keys, and other sensitive information securely. KMS solutions provide centralized management, automated key rotation, and secure storage. Many regulations require the use of KMS for managing encryption keys and certificates. KMS can automatically rotate keys, provide audit trails, and simplify compliance with encryption standards.

HashiCorp Vault is a powerful tool for securely managing sensitive data, such as secrets, API keys, certificates, and encryption keys. Vault is highly versatile and can be used to store and manage secrets for both infrastructure and applications. It supports various use cases, including dynamic secrets, identity-based access, data encryption, and certificate management.


### OPTIONAL

Create a private key for the broker:

`openssl genpkey -algorithm RSA -out kafka-broker.key -aes256`

Generate a certificate signing request (CSR) for the broker:

`openssl req -new -key kafka-broker.key -out kafka-broker.csr`

Generate a self-signed broker certificate (or have it signed by a CA):

`openssl x509 -req -in kafka-broker.csr -signkey kafka-broker.key -out kafka-broker.crt -days 365`

Create a certificate authority (CA) certificate (for mutual authentication if needed):

`openssl genpkey -algorithm RSA -out ca.key`
`openssl req -new -key ca.key -out ca.csr`
`openssl x509 -req -in ca.csr -signkey ca.key -out ca.crt -days 365`

Sign the broker certificate with the CA:

`openssl x509 -req -in kafka-broker.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out kafka-broker.crt -days 365`

(Optional) Generate client certificates (if you're doing mutual authentication):

`openssl genpkey -algorithm RSA -out kafka-client.key`
`openssl req -new -key kafka-client.key -out kafka-client.csr`
`openssl x509 -req -in kafka-client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out kafka-client.crt -days 365`

Now that you have the broker certificate and private key, you can configure the Kafka brokers to use them: In the server.properties file of the Kafka broker, update the following settings to point to the Vault-generated certificates:

```bash
listeners=SSL://0.0.0.0:9093
listener.security.protocol=SSL
ssl.keystore.location=/path/to/kafka-broker.keystore.jks
ssl.keystore.password=broker-keystore-password
ssl.key.password=broker-key-password
ssl.truststore.location=/path/to/kafka-broker.truststore.jks
ssl.truststore.password=broker-truststore-password
ssl.client.auth=required  # Optional, if you want to enforce client certificates
```

The producer (which sends messages to Kafka) needs to authenticate itself to the Kafka broker, usually via SSL/TLS encryption. If the broker requires client authentication (i.e., the broker verifies the producer's identity), the producer will need:

    Keystore: Contains the producer's private key and certificate.

    Truststore: Contains the CA certificates to verify the broker's certificate.

You need a keystore for the broker's private key and certificate, and a truststore for the CA certificate

`keytool -keystore kafka-broker.keystore.jks -alias kafka-broker -keyalg RSA -genkey -validity 3650	`

Then, import the broker certificate and the CA certificate into the keystore and truststore:

`keytool -keystore kafka-broker.keystore.jks -import -alias broker-cert -file kafka-broker.crt`
`keytool -keystore kafka-broker.truststore.jks -import -alias ca-cert -file ca.crt`



		Producer: Requires a keystore (for its certificate) and a truststore (for the broker's CA certificate).

		Consumer: Requires a keystore (if mutual authentication is needed) and a truststore (for the broker's CA certificate).

		Both need the truststore at a minimum to validate the broker’s certificate, but the keystore is needed for client authentication (mutual TLS).

## Difference Between Kafka and RabbitMQ SSL Configuration

While both Kafka and RabbitMQ use SSL/TLS for encryption and mutual authentication, the specific configuration files and tools for storing and managing certificates differ slightly:

    Kafka uses server.properties for broker configuration and may use Java Keystores for storing certificates.

    RabbitMQ uses rabbitmq.conf for SSL/TLS settings and can use PEM files (standard for RabbitMQ) for storing certificates and private keys.

    RabbitMQ Broker needs the CA certificate, server certificate, and private key.

    Clients (Producer/Consumer) need the CA certificate, and if mutual authentication is enabled, the client certificate and private key.

    RabbitMQ uses PEM files for certificates, while Kafka often uses keystore/truststore formats (JKS or PKCS12).

```go

writer, err := kafka.NewProducer(&kafka.ConfigMap{
    "bootstrap.servers":   "localhost:9093",   // SSL port
    "acks":                "all",               // Ensures replication
    "enable.idempotence":  true,                // Enable idempotence for Exactly Once Semantics (EOS)
    "batch.size":          1024 * 1024,        // 1MB batch size
    "linger.ms":           100,                 // Wait for 100ms before sending the batch
    "security.protocol":   "SSL",               // Use SSL (TLS)
    "ssl.ca.location":     "/path/to/ca.crt",   // CA certificate
    "ssl.certificate.location": "/path/to/kafka-client.crt",  // Client certificate
    "ssl.key.location":    "/path/to/kafka-client.key",       // Client private key
    "ssl.key.password":    "your-client-key-password",  // Client key password
})
if err != nil {
    log.Fatalf("Error creating producer: %v", err)
}
defer writer.Close()
```
```go
consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
    "bootstrap.servers":   "localhost:9093",   // SSL port
    "group.id":            "my-consumer-group",
    "auto.offset.reset":   "earliest",
    "security.protocol":   "SSL",               // Use SSL (TLS)
    "ssl.ca.location":     "/path/to/ca.crt",   // CA certificate
    "ssl.certificate.location": "/path/to/kafka-client.crt",  // Client certificate
    "ssl.key.location":    "/path/to/kafka-client.key",       // Client private key
    "ssl.key.password":    "your-client-key-password",  // Client key password
})
if err != nil {
    log.Fatalf("Error creating consumer: %v", err)
}
defer consumer.Close()
```

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// Order represents an order structure.
type Order struct {
	OrderID   string `json:"order_id"`
	Customer  string `json:"customer"`
	Product   string `json:"product"`
	Quantity  int    `json:"quantity"`
	OrderTime string `json:"order_time"`
}

func main() {
	// Kafka producer configuration
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Kafka server address
		"acks":              "all",            // Ensures message replication
		"linger.ms":         100,              // Wait for 100ms to batch messages
		"batch.size":        1024 * 1024,      // 1MB batch size
	})
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer p.Close()

	// Order event to send
	order := Order{
		OrderID:   "12345",
		Customer:  "John Doe",
		Product:   "Laptop",
		Quantity:  1,
		OrderTime: time.Now().Format(time.RFC3339),
	}

	// Marshal order struct to JSON-like string for sending
	orderData := fmt.Sprintf(`{"order_id":"%s","customer":"%s","product":"%s","quantity":%d,"order_time":"%s"}`,
		order.OrderID, order.Customer, order.Product, order.Quantity, order.OrderTime)

	// Send message to Kafka
	topic := "order-events"
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(orderData),
	}, nil)

	if err != nil {
		log.Fatalf("Failed to produce message: %v", err)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000) // Wait for 15 seconds to ensure delivery
	log.Printf("Order event sent: %s", orderData)
}
// To produce messages to a specific Kafka partition using the confluent-kafka-go library, you need to set the partition explicitly in the TopicPartition when sending the message. By default, if you set the partition to kafka.PartitionAny, Kafka will automatically assign the partition for you based on its internal partitioning logic (e.g., using a partitioner).

// Calculate partition from the OrderID (e.g., using a hash function)
	partition := int32(crc32.ChecksumIEEE([]byte(order.OrderID)) % 3) // Assuming 3 partitions, adjust as needed
// Send message to Kafka (specifying partition)
	topic := "order-events"
	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: partition},
		Value:          []byte(orderData),
	}, nil)

```

The producer will collect messages into batches before sending them to Kafka. This `BatchSize` setting defines the maximum size of a single batch.

This `LingerMillis` setting allows the producer to wait for additional messages to fill up the batch before sending it. If a batch hasn't been filled up to the configured BatchSize, the producer will wait for LingerMillis milliseconds before sending it, even if the batch size isn't fully reached.

For `RequiredAcks` Field we have 3 Options:
kafka.NoResponse: No acknowledgment (lowest latency, highest risk).

kafka.WaitForLocal: Acknowledge once the leader replica acknowledges.

kafka.RequireAll: Wait for acknowledgment from all replicas (highest durability).

`IDempotent` setting enables or disables idempotent producers. When idempotency is enabled, Kafka ensures that even if the producer sends the same message multiple times (due to retries), the message is only written once to the Kafka topic.

`Transport` is the TLS configuration. It defines how the connection will be secured with SSL/TLS.
NEVER use InsecureSkipVerify: true in a production environment because it makes your connection vulnerable to man-in-the-middle (MITM) attacks. Always set it to false in production.


```go
package main

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func main() {
	// Kafka consumer configuration
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Kafka server address
		"group.id":          "notification-consumer-group",
		"auto.offset.reset":  "earliest", // Start reading from the earliest offset
	})
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer c.Close()

	// Subscribe to the "order-events" topic
	err = c.SubscribeTopics([]string{"order-events"}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	// Consume messages in a loop
	for {
		msg, err := c.ReadMessage(-1) // Block indefinitely until message is received
		if err != nil {
			// Handle error
			if kafkaError, ok := err.(kafka.Error); ok && kafkaError.Code() == kafka.ErrTimedOut {
				continue
			}
			log.Fatalf("Error while consuming message: %v", err)
		}

		// Print the received order message
		orderData := string(msg.Value)
		log.Printf("Received order event: %s", orderData)

		// Send a notification based on the event (can be an email, SMS, etc.)
		// In this example, we're just printing it to the console
		fmt.Printf("Notification: New Order Received!\nOrder Details: %s\n", orderData)
	}
}

```

## docker-compose.yml file :
```yaml
version: '3'
services:
  zookeeper:
    image: wurstmeister/zookeeper:3.4.6
    ports:
      - "2181:2181"
  kafka:
    image: wurstmeister/kafka:latest
    ports:
      - "9092:9092"
    expose:
      - "9093"
    environment:
      KAFKA_ADVERTISED_LISTENER: INSIDE:9093
      KAFKA_LISTENER: INSIDE
      KAFKA_LISTENER_SECURITY_PROTOCOL: PLAINTEXT
      KAFKA_ZOOKEEPER_CONNECT: zookeeper:2181
      KAFKA_LISTENER_NAME_INSIDE: INSIDE
      KAFKA_LISTENER_PORT: 9093
```