package commands

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/store"
)

type CancelOptions struct {
	SSM   awsclient.SSM
	S3    awsclient.S3
	Store *store.Store

	CommandID         string
	SignalNumber      int
	SignalS3Bucket    string
	SignalS3KeyPrefix string
}

func Cancel(opts *CancelOptions) error {
	resp, err := opts.SSM.ListCommands(&ssm.ListCommandsInput{
		CommandId: aws.String(opts.CommandID),
	})
	if err != nil {
		return err
	}

	if len(resp.Commands) == 0 {
		return errors.New("command is not found")
	}
	status := *resp.Commands[0].Status
	if status != "Pending" && status != "InProgress" {
		return fmt.Errorf("can't cancel the command because its status is %s", status)
	}

	// Get pcommand ID
	r, err := opts.Store.GetCommand(opts.CommandID)
	if err != nil {
		return err
	}
	pcommandID := r.PcommandID

	// Put signal
	payload := map[string]int{
		"signal": opts.SignalNumber,
	}
	j, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s%s.json", opts.SignalS3KeyPrefix, pcommandID)
	log.Printf("DEBUG: putting a signal object to s3://%s/%s", opts.SignalS3Bucket, key)
	_, err = opts.S3.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(j),
		Bucket: aws.String(opts.SignalS3Bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return nil
}
