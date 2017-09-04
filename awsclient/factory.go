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

func (f *Factory) Region() string {
	return *f.sess.Config.Region
}

func (f *Factory) SSM() SSM {
	return SSM(ssm.New(f.sess))
}

func (f *Factory) S3() S3 {
	return S3(s3.New(f.sess))
}
