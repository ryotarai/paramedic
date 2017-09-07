package commands

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type Command struct {
	CommandID  string
	PcommandID string
	Status     string
	Targets    map[string][]string

	OutputLogGroup        string
	OutputLogStreamPrefix string
	SignalS3Bucket        string
	SignalS3Key           string
}

func commandFromSDK(c *ssm.Command, pcommandID string) *Command {
	targets := map[string][]string{}
	for _, t := range c.Targets {
		targets[*t.Key] = aws.StringValueSlice(t.Values)
	}

	return &Command{
		CommandID:             *c.CommandId,
		PcommandID:            pcommandID,
		Status:                *c.Status,
		OutputLogGroup:        *c.Parameters["outputLogGroup"][0],
		OutputLogStreamPrefix: *c.Parameters["outputLogStreamPrefix"][0],
		SignalS3Bucket:        *c.Parameters["signalS3Bucket"][0],
		SignalS3Key:           *c.Parameters["signalS3Key"][0],
		Targets:               targets,
	}
}
