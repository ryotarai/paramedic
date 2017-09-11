package commands

import (
	"github.com/ryotarai/paramedic/documents"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
)

type Command struct {
	CommandID    string
	PcommandID   string
	Status       string
	Targets      map[string][]string
	DocumentName string

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

	doc := documents.ConvertFromSSMName(*c.DocumentName)

	return &Command{
		CommandID:             *c.CommandId,
		PcommandID:            pcommandID,
		Status:                *c.Status,
		OutputLogGroup:        *c.Parameters["outputLogGroup"][0],
		OutputLogStreamPrefix: *c.Parameters["outputLogStreamPrefix"][0],
		SignalS3Bucket:        *c.Parameters["signalS3Bucket"][0],
		SignalS3Key:           *c.Parameters["signalS3Key"][0],
		Targets:               targets,
		DocumentName:          doc,
	}
}

type CommandInvocation struct {
	CommandID    string
	InstanceID   string
	InstanceName string
	Status       string
}

func commandInvocationFromSDK(c *ssm.CommandInvocation) *CommandInvocation {
	return &CommandInvocation{
		CommandID:    *c.CommandId,
		InstanceID:   *c.InstanceId,
		InstanceName: *c.InstanceName,
		Status:       *c.Status,
	}
}
