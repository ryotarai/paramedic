package commands

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/store"
)

type GetOptions struct {
	SSM   awsclient.SSM
	Store *store.Store

	CommandID string
}

func Get(opts *GetOptions) (*Command, error) {
	resp, err := opts.SSM.ListCommands(&ssm.ListCommandsInput{
		CommandId: aws.String(opts.CommandID),
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Commands) == 0 {
		return nil, errors.New("command is not found")
	}

	r, err := opts.Store.GetCommand(opts.CommandID)
	if err != nil {
		return nil, err
	}

	return commandFromSDK(resp.Commands[0], r.PcommandID), nil
}
