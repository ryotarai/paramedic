package awsclient

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type Factory struct {
	sess *session.Session
}

func NewFactory() (*Factory, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	return &Factory{
		sess: sess,
	}, nil
}

func (f *Factory) SSM() *ssm.SSM {
	return ssm.New(f.sess)
}

func (f *Factory) S3() *s3.S3 {
	return s3.New(f.sess)
}
