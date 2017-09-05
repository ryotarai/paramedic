package awsclient

import "github.com/aws/aws-sdk-go/service/s3"

type S3 interface {
	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
}
