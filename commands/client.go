package commands

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/google/uuid"
	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/store"
)

type Client struct {
	SSM   awsclient.SSM
	S3    awsclient.S3
	Store *store.Store
}

// Get a command by ID
func (c *Client) Get(commandID string) (*Command, error) {
	resp, err := c.SSM.ListCommands(&ssm.ListCommandsInput{
		CommandId: aws.String(commandID),
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Commands) == 0 {
		return nil, errors.New("command is not found")
	}

	r, err := c.Store.GetCommand(commandID)
	if err != nil {
		return nil, err
	}

	return commandFromSDK(resp.Commands[0], r.PcommandID), nil
}

// GetInvocations finds command invocations by command ID
func (c *Client) GetInvocations(commandID string) ([]*CommandInvocation, error) {
	invocations := []*CommandInvocation{}

	err := c.SSM.ListCommandInvocationsPages(&ssm.ListCommandInvocationsInput{
		CommandId: aws.String(commandID),
	}, func(resp *ssm.ListCommandInvocationsOutput, last bool) bool {
		for _, i := range resp.CommandInvocations {
			invocations = append(invocations, commandInvocationFromSDK(i))
		}
		return true
	})
	if err != nil {
		return nil, err
	}
	return invocations, nil
}

// Cancel a command by sending signal
func (c *Client) Cancel(command *Command, signal int) error {
	status := command.Status
	if status != "Pending" && status != "InProgress" {
		return fmt.Errorf("can't cancel the command because its status is %s", status)
	}

	// Put signal
	payload := map[string]int{
		"signal": signal,
	}
	j, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Putting a signal object to s3://%s/%s", command.SignalS3Bucket, command.SignalS3Key)
	_, err = c.S3.PutObject(&s3.PutObjectInput{
		Body:   bytes.NewReader(j),
		Bucket: aws.String(command.SignalS3Bucket),
		Key:    aws.String(command.SignalS3Key),
	})
	if err != nil {
		return err
	}

	return nil
}

// GetInstances returns filtered instances
func (c *Client) GetInstances(instanceIDs []string, tags map[string][]string) ([]*Instance, error) {
	filters := []*ssm.InstanceInformationStringFilter{}

	if len(instanceIDs) > 0 {
		filters = append(filters, &ssm.InstanceInformationStringFilter{
			Key:    aws.String("InstanceIds"),
			Values: aws.StringSlice(instanceIDs),
		})
	}

	for k, v := range tags {
		filters = append(filters, &ssm.InstanceInformationStringFilter{
			Key:    aws.String(fmt.Sprintf("tag:%s", k)),
			Values: aws.StringSlice(v),
		})
	}

	instances := []*Instance{}
	err := c.SSM.DescribeInstanceInformationPages(&ssm.DescribeInstanceInformationInput{
		Filters: filters,
	}, func(resp *ssm.DescribeInstanceInformationOutput, last bool) bool {
		for _, info := range resp.InstanceInformationList {
			i := &Instance{
				InstanceID:   *info.InstanceId,
				ComputerName: *info.ComputerName,
			}
			instances = append(instances, i)
		}
		return true
	})
	if err != nil {
		return nil, err
	}

	return instances, nil
}

// GetInstanceIDToNameMap returns a map between instance ID and name
func (c *Client) GetInstanceIDToNameMap() (map[string]string, error) {
	instances, err := c.GetInstances([]string{}, map[string][]string{})
	if err != nil {
		return nil, err
	}

	m := map[string]string{}
	for _, i := range instances {
		m[i.InstanceID] = i.ComputerName
	}
	return m, nil
}

// SendOptions is options for Send
type SendOptions struct {
	DocumentName      string
	InstanceIDs       []string
	Tags              map[string][]string
	MaxConcurrency    string
	MaxErrors         string
	OutputLogGroup    string
	SignalS3Bucket    string
	SignalS3KeyPrefix string
}

// Send a new command
func (c *Client) Send(opts *SendOptions) (*Command, error) {
	pcommandID := uuid.New().String()

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
	resp, err := c.SSM.SendCommand(&ssm.SendCommandInput{
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

	err = c.Store.PutCommand(&store.CommandRecord{
		CommandID:  commandID,
		PcommandID: pcommandID,
	})
	if err != nil {
		return nil, err
	}

	return commandFromSDK(resp.Command, pcommandID), nil
}

// WaitStatus waits a command to be in specified status
func (c *Client) WaitStatus(commandID string, statuses []string) chan *Command {
	interval := 15 * time.Second

	ch := make(chan *Command)

	go func() {
		for {
			log.Printf("[DEBUG] Checking status of command %s", commandID)

			resp, err := c.SSM.ListCommands(&ssm.ListCommandsInput{
				CommandId: aws.String(commandID),
			})
			if err != nil {
				log.Printf("[WARN] %s", err)
			}
			if len(resp.Commands) == 0 {
				log.Printf("[WARN] Command %s is not found", commandID)
			}

			for _, st := range statuses {
				if *resp.Commands[0].Status == st {
					cmd, err := c.Get(commandID)
					if err != nil {
						log.Printf("[WARN] %s", err)
					}

					ch <- cmd
					return
				}
			}

			time.Sleep(interval)
		}
	}()

	return ch
}
