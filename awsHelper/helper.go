package awsHelper

import (
	"bytes"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
	"os"
)

type AwsService interface {
	Upload(bucket, path string, file *os.File) (string, error)
}

type awsObject struct {
	cfg *aws.Config
	svc *s3.S3
}

func New(accessKeyId, secretKey string) AwsService {
	cfg := getConfig(accessKeyId, secretKey)
	return &awsObject{
		cfg: cfg,
		svc: getSvc(cfg),
	}
}

func (receiver *awsObject) Upload(bucket, path string, file *os.File) (string, error) {
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size) // read file content to buffer

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)

	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}

	res, err := receiver.svc.PutObject(params)

	return awsutil.StringValue(res), err
}

func getConfig(accessKeyId, secretKey string) *aws.Config {
	aws_access_key_id := accessKeyId
	aws_secret_access_key := secretKey
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		// handle error
	}
	return aws.NewConfig().WithRegion("eu-central-1").WithCredentials(creds)
}

func getSvc(config *aws.Config) *s3.S3 {
	return s3.New(session.New(), config)
}
