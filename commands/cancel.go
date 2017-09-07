package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/store"
)

type CancelOptions struct {
	SSM   awsclient.SSM
	S3    awsclient.S3
	Store *store.Store

	Command      *Command
	SignalNumber int
}

func Cancel(opts *CancelOptions) error {
	status := opts.Command.Status
	if status != "Pending" && status != "InProgress" {
		return fmt.Errorf("can't cancel the command because its status is %s", status)
	}

	// Put signal
	payload := map[string]int{
		"signal": opts.SignalNumber,
	}
	j, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Putting a signal object to s3://%s/%s", opts.Command.SignalS3Bucket, opts.Command.SignalS3Key)
	_, err = opts.S3.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(j),
		Bucket: aws.String(opts.Command.SignalS3Bucket),
		Key:    aws.String(opts.Command.SignalS3Key),
	})
	if err != nil {
		return err
	}

	return nil
}
