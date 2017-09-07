package commands

import (
	"errors"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/ryotarai/paramedic/awsclient"
)

type Instance struct {
	InstanceID   string
	ComputerName string
}

type GetInstancesOptions struct {
	SSM awsclient.SSM

	InstanceIDs []string
	Tags        map[string][]string
}

func GetInstances(opts *GetInstancesOptions) ([]*Instance, error) {
	if len(opts.InstanceIDs) == 0 && len(opts.Tags) == 0 {
		return nil, errors.New("both instance IDs and tags are not specified")
	}
	filters := []*ssm.InstanceInformationStringFilter{}

	if len(opts.InstanceIDs) > 0 {
		filters = append(filters, &ssm.InstanceInformationStringFilter{
			Key:    aws.String("InstanceIds"),
			Values: aws.StringSlice(opts.InstanceIDs),
		})
	}

	for k, v := range opts.Tags {
		filters = append(filters, &ssm.InstanceInformationStringFilter{
			Key:    aws.String(fmt.Sprintf("tag:%s", k)),
			Values: aws.StringSlice(v),
		})
	}

	instances := []*Instance{}
	err := opts.SSM.DescribeInstanceInformationPages(&ssm.DescribeInstanceInformationInput{
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
