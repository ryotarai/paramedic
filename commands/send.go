package commands

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/google/uuid"
	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/store"
)

type SendOptions struct {
	SSM   awsclient.SSM
	Store *store.Store

	DocumentName      string
	InstanceIDs       []string
	Tags              map[string][]string
	MaxConcurrency    string
	MaxErrors         string
	OutputLogGroup    string
	SignalS3Bucket    string
	SignalS3KeyPrefix string
}

func Send(opts *SendOptions) (*Command, error) {
	pcommandID := generatePcommandID()

	targets := []*ssm.Target{}
	if len(opts.InstanceIDs) > 0 {
		targets = append(targets, &ssm.Target{
			Key:    aws.String("InstanceIds"),
			Values: aws.StringSlice(opts.InstanceIDs),
		})
	}

	for k, v := range opts.Tags {
		targets = append(targets, &ssm.Target{
			Key:    aws.String(fmt.Sprintf("tag:%s", k)),
			Values: aws.StringSlice(v),
		})
	}

	// TODO: write output to S3
	resp, err := opts.SSM.SendCommand(&ssm.SendCommandInput{
		DocumentName:   aws.String(opts.DocumentName),
		Targets:        targets,
		MaxConcurrency: aws.String(opts.MaxConcurrency),
		MaxErrors:      aws.String(opts.MaxErrors),
		Parameters: map[string][]*string{
			"outputLogGroup":        []*string{aws.String(opts.OutputLogGroup)},
			"outputLogStreamPrefix": []*string{aws.String(fmt.Sprintf("%s/", pcommandID))},
			"signalS3Bucket":        []*string{aws.String(opts.SignalS3Bucket)},
			"signalS3Key":           []*string{aws.String(fmt.Sprintf("%s%s.json", opts.SignalS3KeyPrefix, pcommandID))},
		},
	})
	if err != nil {
		return nil, err
	}

	commandID := *resp.Command.CommandId

	err = opts.Store.PutCommand(&store.CommandRecord{
		CommandID:  commandID,
		PcommandID: pcommandID,
	})
	if err != nil {
		return nil, err
	}

	return commandFromSDK(resp.Command, pcommandID), nil
}

func generatePcommandID() string {
	return uuid.New().String()
}
