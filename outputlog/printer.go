package outputlog

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/ryotarai/paramedic/awsclient"
)

type Printer struct {
	CloudWatchLogs awsclient.CloudWatchLogs
	Interval       time.Duration
	LogGroupName   string
	LogStreamName  string
	StartFromHead  bool
}

func (p *Printer) Start() {
	log.Printf("DEBUG: starting to print logs of %s/%s", p.LogGroupName, p.LogStreamName)
	parts := strings.Split(p.LogStreamName, "/")
	instanceID := parts[len(parts)-1]

	var nextToken string

	for {
		log.Printf("DEBUG: getting logs of %s/%s", p.LogGroupName, p.LogStreamName)
		input := &cloudwatchlogs.GetLogEventsInput{
			LogGroupName:  aws.String(p.LogGroupName),
			LogStreamName: aws.String(p.LogStreamName),
			StartFromHead: aws.Bool(p.StartFromHead),
		}
		if nextToken != "" {
			input.NextToken = aws.String(nextToken)
		}

		err := p.CloudWatchLogs.GetLogEventsPages(input, func(resp *cloudwatchlogs.GetLogEventsOutput, last bool) bool {
			if len(resp.Events) == 0 {
				nextToken = *resp.NextForwardToken
				return false
			}
			for _, e := range resp.Events {
				t := time.Unix(0, (*e.Timestamp)*1000*1000)
				fmt.Printf("%s | %s | %s\n", instanceID, t.Format(time.RFC3339), *e.Message)
			}
			return true
		})

		if err != nil {
			log.Printf("ERROR: %s", err)
		}

		if p.Interval == time.Duration(0) {
			break
		}

		time.Sleep(p.Interval)
	}
}
