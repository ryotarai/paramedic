package cmd

import (
	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/commands"
	"github.com/ryotarai/paramedic/documents"
	"github.com/ryotarai/paramedic/store"
)

func newCommandsClient(f *awsclient.Factory) (*commands.Client, error) {
	return &commands.Client{
		SSM:   f.SSM(),
		S3:    f.S3(),
		Store: store.New(f.DynamoDB()),
	}, nil
}

func newDocumentsClient(f *awsclient.Factory, bucket, keyPrefix string) (*documents.Client, error) {
	return &documents.Client{
		SSM:               f.SSM(),
		S3:                f.S3(),
		ScriptS3Bucket:    bucket,
		ScriptS3KeyPrefix: keyPrefix,
	}, nil
}
