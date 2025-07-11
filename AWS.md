

# What are the main services provided by AWS?
- Compute: EC2, Lambda, Elastic Beanstalk
- Storage: S3, EBS, Glacier
- Database: RDS, DynamoDB, Aurora
- Networking: VPC, CloudFront, Route53
- Security: IAM, KMS, Shield, WAF

# Services:
- A Region is a geographical location where AWS has multiple data centers (Availability Zones). Each region is isolated from others 

- An Availability Zone (AZ) is one or more physically separate data centers within a region. Each region has 2 or more Availability Zones.

- Amazon VPC (Virtual Private Cloud) is a logically isolated network within the AWS cloud. We can define our own IP address ranges, create subnets, configure route tables, and attach internet gateways, NAT gateways, and VPN connections.

- A Security Group in AWS acts as a virtual firewall that controls both inbound and outbound traffic to EC2 instances and other resources. For example, if I allow incoming traffic on port 80 for HTTP, the outgoing response is automatically allowed — that's what we mean by stateful. I typically use security groups to restrict access to only necessary IP addresses or other AWS services. For instance, I might open port 22 only to my corporate IP for SSH, or restrict a database instance to accept traffic only from the web server’s security group, not from the public internet.

- AWS IAM, or Identity and Access Management, is the core security service that enables you to manage access to AWS resources securely. I use IAM to define who can access what, and under what conditions. For example, I create IAM users for individual team members, groups for roles like DevOps or Developers, and attach policies that define their permissions. I use roles for granting temporary permissions — such as allowing EC2 instances to access S3 buckets, or enabling Lambda to call DynamoDB.

- Amazon EC2 (Elastic Compute Cloud) is a virtual server that allows users to run applications on AWS’s infrastructure. You can choose the OS, instance type, storage, and networking.

- AWS Lambda is a serverless compute service that allows you to run code without provisioning or managing servers. You simply upload your code, specify the trigger (such as an S3 event, SNS message, or API Gateway request), and Lambda automatically scales the execution based on incoming requests. AWS Lambda charges based on the number of requests and the compute time consumed (measured in GB-seconds).  And supports multiple languages like Go, Node.js, Python, Java, and C#.

- Amazon EBS is a block-level storage service designed for use with Amazon EC2 instances. It's like a virtual hard drive that can be attached to EC2 instances to store data, operating system files, applications, and more. Unlike traditional hard drives, EBS volumes are persistent, meaning data remains intact even when the EC2 instance is stopped or terminated (unless explicitly deleted).

- Amazon S3 is a highly scalable object storage service that allows you to store and retrieve an unlimited amount of data. It’s perfect for a wide range of use cases, from backup storage to big data analytics and media hosting. You can organize your data using buckets (containers for objects), and each object is uniquely identified by a key within a bucket. I’ve used EBS volumes to store operating system files, databases, and application data that requires fast access and performance. For example, if you're running a database on EC2, you’d want to store that data on EBS. On the other hand, I’ve used S3 for storing static assets, backups, and logs because of its durability and cost-effectiveness for large-scale storage. I also use EBS snapshots to back up EC2 instance volumes and then store those backups in S3 to combine both storage types for different use cases.

- Amazon SQS is a fully managed message queuing service that helps decouple and scale microservices, distributed systems, and serverless applications. It allows you to send, store, and receive messages between software components at any volume, without losing messages. SQS supports both standard queues (the order in which messages are received is not guaranteed. So, messages could be processed out of order, every message will be delivered at least once, but there might be duplicates. ) and FIFO queues (ensures that the messages are processed in the exact order they were sent, each message will be delivered exactly one time, with no duplicates).

- Amazon SNS is a fully managed pub/sub messaging service that allows you to send messages to multiple subscribers simultaneously. It supports multiple protocols, including SMS, email, HTTP(S), and even Lambda functions. SNS enables the broadcasting of messages to a wide range of endpoints, making it ideal for use cases like application alerts, event notifications, and push notifications."

- Auto Scaling automatically adjusts the number of EC2 instances in a group based on traffic or load, helping ensure application availability and cost optimization.

- AWS uses Elastic Load Balancers (ELB) to automatically distribute incoming application traffic across multiple targets (like EC2). There are Application, Network, and Gateway Load Balancers for different use cases.

- CloudWatch is used for monitoring AWS resources and applications. It collects logs, metrics, and allows you to set alarms or automate actions.


I used S3 to store music tracks, metadata, and other assets (like album covers) uploaded. When a Label uploads a new song and artwork and metadata, the file gets stored in an S3 bucket. This also triggered events for further processing using Lambda.

A Lambda function would automatically get triggered to generate different versions of the track (e.g., high-quality Flac, low-quality mp3, etc.) and store the processed files back in S3. Lambda helps here because it is serverless, automatically scaling as needed without the overhead of managing servers.

When Lambda finishes processing the song, it places a message in an SQS queue, signaling that the song is ready for distribution. Another service picks up the message from the queue to update the song's metadata in the database.

When the music processing is completed, an SNS message is sent to the artist to notify them that their track is live and available. 

| **Service** | **Attachments/Configurations**                                                                                     |
|-------------|---------------------------------------------------------------------------------------------------------------------|
| **S3**      | - Bucket Policies/ACLs <br> - IAM Role for access <br> - Event Notifications (Lambda, SNS, SQS)                     |
| **SQS**     | - IAM Role/Policy for access <br> - Dead Letter Queue (optional) <br> - Queue Permissions                           |
| **SNS**     | - Topic Policy for permissions <br> - Subscriptions (Lambda, SQS, Email, HTTP) <br> - IAM Role for publishing       |
| **Lambda**  | - Execution Role (IAM role with permissions) <br> - Event Sources (S3, SQS, SNS, API Gateway) <br> - Environment Variables |
# Build Go Binaries
```bash
GOOS=linux GOARCH=amd64 go build -o user main.go
```
# PORTS
80: Http
443: Https
22: SSH

# SSH to EC2 instance
ssh -i my.pem ec2-user@124:23:21:12

# SCP to EC2 instance
scp -i my.pem user ec2-user@124:23:21:12


# S3 Operations

```go
cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
client := s3.NewFromConfig(cfg)
file, err := os.Open(filePath)
input := &s3.PutObjectInput{
		Bucket: &bucket,
		Key:    &key,
		Body:   file,
	}
	_, err = client.PutObject(context.TODO(), input)
    fmt.Println("Upload successful!")
    output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &key,
	})
    f, err := os.Create(localFile)
    _, err = io.Copy(f, output.Body)
    // _, err = f.ReadFrom(output.Body)
    fmt.Println("Download successful!")
```
Multipart upload is useful for uploading large files because it breaks the file into smaller parts and uploads them independently. After all parts are uploaded, they are assembled into the final object in S3.
```go
createResp, err := client.CreateMultipartUpload(ctx, &s3.CreateMultipartUploadInput{
	Bucket: aws.String(bucketName),
	Key:    aws.String(key),
})
```
client.CreateMultipartUpload: This method starts a multipart upload. It returns a response (createResp) that includes an UploadId that will be used in the subsequent part uploads.
```go
file, err := os.Open("path/to/large/file")
var completedParts []s3types.CompletedPart // This slice holds information about each successfully uploaded part (ETag and part number). It will be used when completing the multipart upload.
var partNumber int32 = 1

for {
	partBuffer := make([]byte, partSize)// Creates a buffer (partBuffer) to hold each chunk of the file. The buffer size is partSize, which is 5 MB.
	n, err := file.Read(partBuffer)
	if err != nil && n == 0 {
		break
	}

	uploadPartResp, err := client.UploadPart(ctx, &s3.UploadPartInput{
		Bucket:     aws.String(bucketName),
		Key:        aws.String(key),
		UploadId:   createResp.UploadId,
		PartNumber: partNumber,
		Body:       bytes.NewReader(partBuffer[:n]), //  It takes the part data from the buffer (partBuffer[:n]), which is the portion of the file being uploaded.
	})
	if err != nil {
		log.Fatalf("Upload part %d failed: %v", partNumber, err)
	}
    // If the part upload is successful, it appends the ETag and PartNumber to the completedParts slice, which will be used to finalize the upload.
	completedParts = append(completedParts, s3types.CompletedPart{
		ETag:       uploadPartResp.ETag,
		PartNumber: partNumber,
	})
	partNumber++
}
```
Once all parts are uploaded, CompleteMultipartUpload is called to combine the parts into a single file in S3.
```go
_, err = client.CompleteMultipartUpload(ctx, &s3.CompleteMultipartUploadInput{
	Bucket:   aws.String(bucketName),
	Key:      aws.String(key),
	UploadId: createResp.UploadId,
	MultipartUpload: &s3types.CompletedMultipartUpload{
		Parts: completedParts,
	},
})
```
# Lambda
Want to trigger the Lambda through an HTTP request so need AWS API Gateway
Lambda doesn’t natively expose an HTTP endpoint. API Gateway acts as the front door, letting users or apps call your Lambda via standard HTTP methods (GET, POST, etc.).
The handler in an AWS Lambda function is essentially the entry point—the method that gets executed when your function is invoked.
```go
func handler(req events.APIGatewayProxyRequest) events.APIGatewayProxyResponse {
    fileName:= request.QueryStringParameters["file"] // Query parameter for file name
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
    client := s3.NewFromConfig(cfg)
    output, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &fileName,
	})
    fileContent, err := ioutil.ReadAll(output.Body)
    return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(fileContent), // File content as a string
	},
}
func main() {
	lambda.Start(handler)
}
```
To deploy this Lambda function, we need to: Build the Go executable and Create the deployment package: 
```bash
GOOS=linux GOARCH=amd64 go build -o main
zip function.zip main 
```


Upload the function.zip file to Lambda Console. 

Set role to allow the Lambda function to access S3. 

Go to the IAM Console: Attach a policy with s3:GetObject permission to the Lambda role. 

Go to the API Gateway Console and create a new API. Choose HTTP API (simpler) or REST API (more features). Create a route (e.g., /download). Set Lambda function as the integration type and choose the Lambda function you just created.

curl "https://xyz12345.execute-api.us-west-2.amazonaws.com/dev/download?file=my-file.txt"

# SQS

`Amazon SQS Queue Types`
- Standard Queue handles massive volumes of messages.At-least-once delivery — duplicates may occur. No guaranteed order — messages may arrive out of sequence. Use Cases: Background jobs, real-time data processing, microservice communication.
- FIFO Queue (First-In-First-Out) Exactly-once delivery — no duplicates. Guaranteed order — messages are processed as sent.Message grouping — groups can be processed in parallel, maintaining order within each. Use Cases: Financial transactions, inventory updates, workflow orchestration.


# SNS
To allow SNS to send messages to the SQS queue, you need to add an SQS Queue Policy that grants sns.amazonaws.com permission to send messages.

`Two types of topics to suit different messaging needs: `
- Standard Topic High throughput — supports millions of messages per second. At-least-once delivery — messages may be duplicated. No guaranteed order — messages can arrive out of sequence. Supports multiple protocols — SQS, Lambda, HTTP/S, SMS, email, mobile push. Use Cases: Real-time alerts, fan-out messaging, background processing.
- FIFO Topic (First-In-First-Out) Strict ordering — messages are delivered exactly as published. Exactly-once delivery — no duplicates. Limited subscriptions — only supports FIFO SQS queues. Lower throughput — up to 300 messages/sec or 10 MB/sec. Use Cases: Financial transactions, inventory updates, workflows needing order.



```go
package main

import (
    "fmt"
    "net/http"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/sqs"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
    file, header, _ := r.FormFile("file")
    defer file.Close()

    sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
    s3Client := s3.New(sess)

    _, err := s3Client.PutObject(&s3.PutObjectInput{
        Bucket: aws.String("input-bucket"),
        Key:    aws.String(header.Filename),
        Body:   file,
    })
    if err != nil {
        http.Error(w, "S3 upload failed", 500)
        return
    }

    sqsClient := sqs.New(sess)
    _, err = sqsClient.SendMessage(&sqs.SendMessageInput{
        QueueUrl:    aws.String("https://sqs.us-east-1.amazonaws.com/123456789012/my-queue"),
        MessageBody: aws.String(fmt.Sprintf(`{"filename":"%s"}`, header.Filename)),
    })
    if err != nil {
        http.Error(w, "SQS send failed", 500)
        return
    }

    w.Write([]byte("Upload successful"))
}

func main() {
    http.HandleFunc("/upload", uploadHandler)
    http.ListenAndServe(":8080", nil)
}
```


```go
package main

import (
    "context"
    "encoding/json"
    "os/exec"
    "strings"
    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
)

type Message struct {
    Filename string `json:"filename"`

}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
    sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
    s3Client := s3.New(sess)

    for _, record := range sqsEvent.Records {
        var msg Message
        json.Unmarshal([]byte(record.Body), &msg)
        input := "/tmp/" + msg.Filename
        // Download from input bucket
        s3Client.DownloadFile("input-bucket", msg.Filename, input)

        // Convert using FFmpeg
		base := strings.TrimSuffix(msg.Filename, ".mp3")
        flac := "/tmp/" + base + ".flac"
        wav := "/tmp/" + base + ".wav"
        exec.Command("/opt/bin/ffmpeg", "-i", input, flac).Run()
        exec.Command("/opt/bin/ffmpeg", "-i", input, wav).Run()

        // Upload to output bucket
        s3Client.UploadFile("output-bucket", base+".flac", flac)
        s3Client.UploadFile("output-bucket", base+".wav", wav)
    }
	sess := session.Must(session.NewSession())
    snsClient := sns.New(sess)

    _, err := snsClient.Publish(&sns.PublishInput{
        TopicArn: aws.String("arn:aws:sns:us-east-1:123456789012:AudioNotify"),
        Message:  aws.String("Your audio files have been converted and uploaded."),
        Subject:  aws.String("Audio Conversion Complete"),
    })
    return nil
}

func main() {
    lambda.Start(handler)
}
```
```go
// Client
import (
    "context"
    "encoding/json"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/lambda"
    "github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

func callLambda(functionName string, payload interface{}) (string, error) {
    cfg, _ := config.LoadDefaultConfig(context.TODO())
    client := lambda.NewFromConfig(cfg)

    body, _ := json.Marshal(payload)
    input := &lambda.InvokeInput{
        FunctionName:   &functionName,
        Payload:        body,
        InvocationType: types.InvocationTypeRequestResponse,
    }

    result, err := client.Invoke(context.TODO(), input)
    if err != nil {
        return "", err
    }

    return string(result.Payload), nil
}
```
```go
snsClient.Subscribe(&sns.SubscribeInput{
  TopicArn: topic.TopicArn,
  Protocol: aws.String("email"),
  Endpoint: aws.String("user@example.com"),
})
```