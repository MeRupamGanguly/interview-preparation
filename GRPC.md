- Protocol Buffers are a language-neutral, platform-neutral serialization mechanism developed by Google. It is used by gRPC to define service contracts and message formats. 

- Unary RPC – Client sends one request, gets one response. Server streaming – Client sends one request, server sends a stream of responses. Client streaming – Client sends a stream of requests, server responds once. Bi-directional streaming – Both client and server send a stream of messages.

- gRPC uses HTTP/2 and binary serialization (Protocol Buffers) for high performance and efficiency, It supports bi-directional streaming.
- REST relies on HTTP/1.1 and text-based formats like JSON, offering simplicity and ease of use for web applications. REST is stateless and leverages standard HTTP methods, making it widely adopted and easy to debug. 


```proto
syntax="proto3";

package gogen;

option go_package = "/gogen";

message TickerRequest {
    string symbol = 1;
}

message TickerResponse {
    string symbol = 1;
    double ltp = 2;
    string timestamp = 3; 
}

service TickerStreamService{
    rpc TickerStream (stream TickerRequest)returns(stream TickerResponse);
}
```

option go_package = "github.com/MeRupamGanguly/rupamic" this line decides where generated file will store.

Steps to gen grpc-go Codes:
```bash
docker run -it golang:1.23.6 /bin/bash
apt update
apt install protobuf-compiler
go version
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
docker cp ticker/ e3a5b0647829:/ticker/ #New Terminal from Host
protoc --proto_path=/ticker/domain --go_out=/ticker/domain --go-grpc_out=/ticker/domain /ticker/domain/ticker.proto
```
Second Create Dockerfile for generating grpc-go codes:
```dockerfile
FROM golang:1.23.6
RUN apt update
COPY /ticker /ticker
RUN apt install protobuf-compiler -y
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
RUN mkdir -p /ticker/domain/gogen
RUN protoc --proto_path=/ticker/domain --go_out=/ticker/domain --go-grpc_out=/ticker/domain /ticker/domain/ticker.proto
```
Build from Dockerfile
```bash
docker build -t grpcgen -f dockerfile.ticker .
docker run -it grpcgen
docker cp 52509fb095ed:/ticker/github.com/MeRupamGanguly/rupamic/ .
```

- In gRPC we can authenticate and secure connections by using SSL and JWT Token.   SL/TLS ensures that the communication between the client and server is encrypted. With token-based authentication, the client must include a token (like JWT) in the request, which proves their identity and permissions.

- In SSL/TLS authentication we need Certificate and Private Key.

- To generate a certificate (cert) and private key (key) for SSL/TLS authentication in Go (or in any other application), we can use the openssl tool. 
To generate a private key, we can use the following command: `openssl genpkey -algorithm RSA -out server.key -aes256`
To generate Self-Signed Certificate : `openssl req -new -x509 -key server.key -out server.crt -days 365`
Creating a Certificate Signed by a Trusted CA:
- Create a Private Key
- Once we have the private key, we can generate a CSR. This CSR contains information about the server, and it will be submitted to the CA for signing. we can use the following command: `openssl req -new -key server.key -out server.csr` You will be prompted for details like Common Name (CN), Organization, and Location. The most important field is CN (Common Name), which should match the fully qualified domain name (FQDN) of our server (e.g., example.com or ourserver.com).
- You need to submit the server.csr file to a CA (trusted organization) for signing. If we are using a public CA, like Let's Encrypt, GoDaddy, or DigiCert, they will provide a process for submitting the CSR and issuing a signed certificate.
- After the CA signs the CSR, we will receive a signed certificate (server.crt). This certificate is now trusted (assuming it's signed by a trusted CA).
- After obtaining the signed certificate, we can verify that the certificate matches the private key and is properly signed by the CA:

```go
// Load server certificate and private key for TLS
creds, err := credentials.NewServerTLSFromFile("server.crt", "server.key")
```
```go
// Create a gRPC server with TLS credentials
grpcServer := grpc.NewServer(grpc.Creds(creds))
```

```go
// Create a gRPC server with interceptor for JWT validation in main
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(func(
			ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			// Perform authentication
			if _, err := authenticate(ctx); err != nil {
				return nil, grpc.Errorf(grpc.Code(err), err.Error())
			}
			// Call the handler to proceed with the request
			return handler(ctx, req)
		}),
	)
```

The client must also use the server's public certificate to verify the server’s identity and establish a secure connection.

```go
// Load the server certificate
creds, err := credentials.NewClientTLSFromFile("server.crt", "")
// Create a gRPC client with TLS credentials
conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(creds))
```

```proto
syntax = "proto3";

package chat;

service Chat {

  rpc ServerStream (Message) returns (stream Message);


  rpc ClientStream (stream Message) returns (Message);


  rpc StreamChat (stream Message) returns (stream Message);
}


message Message {
  string sender = 1;
  string content = 2;
}

```
- The Chat service contains three RPC methods.
- The Message message type contains two fields: sender (to indicate who sent the message) and content (the message text).

To generate the Go code from this .proto file, use the following command:
```bash
protoc --go_out=. --go-grpc_out=. chat.proto

protoc --proto_path=/ticker/domain --go_out=/ticker/domain --go-grpc_out=/ticker/domain /ticker/domain/ticker.proto
```
```go
package main

import (
    "io"
    "log"
    "net"
    "time"

    "golang.org/x/net/context"
    "google.golang.org/grpc"
    pb "path/to/our/chat" // Adjust the import path
)

// Server struct
type server struct {
    pb.UnimplementedChatServer
}

// Server-side streaming method
func (s *server) ServerStream(req *pb.Message, stream pb.Chat_ServerStreamServer) error {
    for i := 0; i < 5; i++ {
        stream.Send(&pb.Message{Sender: "Server", Content: req.Content + " " + time.Now().String()})
        time.Sleep(time.Second)
    }
    return nil
}

// Client-side streaming method
func (s *server) ClientStream(stream pb.Chat_ClientStreamServer) error {
    var messages []string
    for {
        msg, err := stream.Recv()
        if err == io.EOF {
            break
        }
        messages = append(messages, msg.Content)
    }
    stream.SendAndClose(&pb.Message{Sender: "Server", Content: "Received: " + fmt.Sprint(messages)})
    return nil
}

// Bidirectional streaming method
func (s *server) StreamChat(stream pb.Chat_StreamChatServer) error {
    for {
        msg, err := stream.Recv()
        if err == io.EOF {
            break
        }
        stream.Send(&pb.Message{Sender: "Server", Content: "Echo: " + msg.Content})
    }
    return nil
}

func main() {
    lis, _ := net.Listen("tcp", ":50051")
    s := grpc.NewServer()
    pb.RegisterChatServer(s, &server{})
    go s.Serve(lis)

    conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
    client := pb.NewChatClient(conn)

    // Client-side streaming
    clientStream, _ := client.ClientStream(context.Background())
    for i := 0; i < 3; i++ {
        clientStream.Send(&pb.Message{Content: "Client message " + fmt.Sprint(i)})
    }
    clientStream.CloseSend()

    // Server-side streaming
    serverStream, _ := client.ServerStream(context.Background(), &pb.Message{Content: "Hello!"})
    for {
        msg, err := serverStream.Recv()
        if err != nil {
            break
        }
        log.Println(msg.Content)
    }

    // Bidirectional streaming
    bidiStream, _ := client.StreamChat(context.Background())
    go func() {
        for i := 0; i < 3; i++ {
            bidiStream.Send(&pb.Message{Content: "Bidirectional message " + fmt.Sprint(i)})
            time.Sleep(time.Second)
        }
        bidiStream.CloseSend()
    }()
    for {
        msg, err := bidiStream.Recv()
        if err != nil {
            break
        }
        log.Println(msg.Content)
    }
}
```
Summary of Steps

- Define the Server: A server struct implements the ChatServer interface, with three methods for server-side streaming, client-side streaming, and bidirectional streaming.
- Implement Streaming Methods: The ServerStream method sends multiple messages back to the client; the ClientStream method collects messages from the client and sends a summary; and StreamChat handles both incoming and outgoing messages simultaneously.
- Set Up the gRPC Server: The main function creates a listener on port 50051, registers the server, and starts serving in a separate goroutine.
- Create and Use a gRPC Client: The client connects to the server and demonstrates client-side streaming by sending messages, server-side streaming by receiving multiple messages, and bidirectional streaming in a separate goroutine.
- Run the Server and Client: The entire setup allows for real-time message exchange, showcasing the flexibility of gRPC for various streaming patterns within a single file.

# POINTS
- gRPC uses status codes and metadata to handle errors. For example: grpc.StatusCode.INVALID_ARGUMENT, grpc.StatusCode.UNAUTHENTICATED

- A server-streaming RPC in gRPC allows the server to send multiple messages in response to a single client request. After the client sends a request, the server can stream back a series of responses until the stream is closed. The client can receive messages continuously during the RPC.

- If the client or server does not properly manage the streaming connection (e.g., by not reading or sending messages in the expected order), the stream may hang, time out, or fail. The connection could also be prematurely closed, causing communication failures. Proper error handling and stream management are necessary to ensure the stability of the streaming RPC.

- Clients can specify a deadline or timeout for each RPC call. If the server does not respond within that time, the call fails with a DEADLINE_EXCEEDED error.

- Synchronous (blocking): Client waits for a response. Asynchronous (non-blocking): Client continues execution and handles response via callbacks or futures.

- Use versioned package names in .proto files. Maintain backward compatibility by adding new fields with optional/default values.

- gRPC supports built-in client-side load balancing. Integrates with DNS-based or custom name resolvers.

- Kubernetes assigns a DNS name to each Service (e.g., my-service.default.svc.cluster.local). gRPC clients can connect to these DNS names. Load balancing is handled by Kubernetes using kube-proxy (usually Round Robin).

            You define a Deployment for our gRPC server.

            You expose it via a Service (usually ClusterIP or LoadBalancer).

            The gRPC client calls the server using the service name.

            grpc.Dial("user-service.default.svc.cluster.local:50051", ...)

-  Network hiccups, transient server errors, or load balancer issues can cause temporary failures—retries give the system a second (or third) chance to succeed without user impact.  gRPC retries can be implemented in two main ways: using client-side service configuration or manual retry logic via interceptors. With service config (supported in Go and Java), we define a retryPolicy in the service’s config JSON, specifying rules like max attempts, backoff, and retryable status codes—this is automatic and built into the gRPC client. On the other hand, interceptors allow us to manually code retry logic by wrapping client calls and reissuing them based on custom conditions. While service config is simpler and declarative, interceptors offer more control and flexibility. For basic retry needs, service config is recommended when available. 
