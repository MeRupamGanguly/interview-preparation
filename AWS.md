

# What are the main services provided by AWS?
- Compute: EC2, Lambda, Elastic Beanstalk
- Storage: S3, EBS, Glacier
- Database: RDS, DynamoDB, Aurora
- Networking: VPC, CloudFront, Route53
- Security: IAM, KMS, Shield, WAF

# Services:
- A Region is a geographical location where AWS has multiple data centers (Availability Zones). Each region is isolated from others for fault tolerance, latency control.

- An Availability Zone (AZ) is one or more physically separate data centers within a region, connected with low-latency links. Each region has 2 or more Availability Zones.

- A Security Group in AWS acts as a stateful virtual firewall that controls both inbound and outbound traffic to EC2 instances and other resources. For example, if I allow incoming traffic on port 80 for HTTP, the outgoing response is automatically allowed — that's what we mean by stateful. I typically use security groups to restrict access to only necessary IP addresses or other AWS services. For instance, I might open port 22 only to my corporate IP for SSH, or restrict a database instance to accept traffic only from the web server’s security group, not from the public internet.

- Amazon Route 53 is AWS’s highly available and scalable DNS service - translating domain names into IP addresses. But beyond basic DNS resolution, Route 53 also integrates tightly with AWS to support advanced routing policies, health checks, and even domain registration. We can chose latency based routing where we can routing to direct users to the nearest AWS region for faster performance. If a web server goes down, Route 53 can automatically remove it from DNS responses, ensuring users are routed only to healthy resources. We can used geolocation routing to direct users from different countries to region-specific content, helping meet compliance and improve localization.

- AWS IAM, or Identity and Access Management, is the core security service that enables you to manage access to AWS resources securely. I use IAM to define who can access what, and under what conditions. For example, I create IAM users for individual team members, groups for roles like DevOps or Developers, and attach policies that define their permissions. I use roles for granting temporary permissions — such as allowing EC2 instances to access S3 buckets, or enabling Lambda to call DynamoDB.

- Amazon VPC (Virtual Private Cloud) is a logically isolated network within the AWS cloud. We can define our own IP address ranges, create subnets, configure route tables, and attach internet gateways, NAT gateways, and VPN connections.

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

```go
cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
client := sqs.NewFromConfig(cfg)
queueUrl := "https://sqs.us-east-1.amazonaws.com/123456789012/your-queue"
message := `{"bucket": "your-s3-bucket", "key": "largefile.dat"}`
_, err = client.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueUrl),
		MessageBody: aws.String(message),
	})
```

# SNS
To allow SNS to send messages to the SQS queue, you need to add an SQS Queue Policy that grants sns.amazonaws.com permission to send messages.
```go
cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
sqsClient := sqs.NewFromConfig(cfg)
// Create the SQS Queue
	sqsResp, err := sqsClient.CreateQueue(context.TODO(), &sqs.CreateQueueInput{
		QueueName: aws.String("MyQueue"),
	})
snsClient := sns.NewFromConfig(cfg)
resp, err := snsClient.CreateTopic(context.TODO(), &sns.CreateTopicInput{
		Name: aws.String("MyTopic"),
	})
// Set the Topic ARN (replace with your topic ARN)
	topicArn := "arn:aws:sns:us-east-1:123456789012:MyTopic"
// Publish a message to the SNS topic
	_, err = snsClient.Publish(context.TODO(), &sns.PublishInput{
		TopicArn: aws.String(topicArn),
		Message:  aws.String("This is a test message from Go!"),
	})
// Set the Topic ARN (replace with your topic ARN)
	topicArn := "arn:aws:sns:us-east-1:123456789012:MyTopic"

	// Set the SQS Queue ARN (replace with your SQS ARN)
	sqsQueueArn := "arn:aws:sqs:us-east-1:123456789012:MyQueue"
// Subscribe the SQS Queue to the SNS Topic
	_, err = snsClient.Subscribe(context.TODO(), &sns.SubscribeInput{
		Protocol: aws.String("sqs"),
		Endpoint: aws.String(sqsQueueArn),
		TopicArn: aws.String(topicArn),
	})
```