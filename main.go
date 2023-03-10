package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"golang-s3/awsHelper"
	"os"
)

func main() {

	aws_access_key_id := "access_key"
	aws_secret_access_key := "secret_key"

	awsInstance := awsHelper.New(aws_access_key_id, aws_secret_access_key)

	file, err := os.Open("rft.jpg")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer file.Close()

	path := "/media/" + file.Name()
	bucket := "myfirstbucketrafet"
	resp, err := awsInstance.Upload(bucket, path, file)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("response %s", awsutil.StringValue(resp))
}
