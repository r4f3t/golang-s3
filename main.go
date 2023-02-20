package main

import (
	"bytes"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"golang-s3/awsHelper"
	"net/http"
	"os"
)

func main() {

	aws_access_key_id := "AKIASGXB2POGYQQUXUUQ"
	aws_secret_access_key := "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"

	awsInstance := awsHelper.New(aws_access_key_id, aws_secret_access_key)

	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	_, err := creds.Get()
	if err != nil {
		// handle error
	}
	cfg := aws.NewConfig().WithRegion("eu-central-1").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	file, err := os.Open("rft.jpg")
	if err != nil {
		// handle error
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size) // read file content to buffer

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := "/media/" + file.Name()
	params := &s3.PutObjectInput{
		Bucket:        aws.String("myfirstbucketrafet"),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	resp, err := svc.PutObject(params)
	if err != nil {
		// handle error
	}
	fmt.Printf("response %s", awsutil.StringValue(resp))
}
