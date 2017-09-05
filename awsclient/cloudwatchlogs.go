package awsclient

import (
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
)

type CloudWatchLogs interface {
	DescribeLogStreamsPages(*cloudwatchlogs.DescribeLogStreamsInput, func(*cloudwatchlogs.DescribeLogStreamsOutput, bool) bool) error
	GetLogEventsPages(*cloudwatchlogs.GetLogEventsInput, func(*cloudwatchlogs.GetLogEventsOutput, bool) bool) error
}
