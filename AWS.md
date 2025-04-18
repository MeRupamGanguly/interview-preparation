
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