package outputlog

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/ryotarai/paramedic/awsclient"
)

type Watcher struct {
	CloudWatchLogs      awsclient.CloudWatchLogs
	Interval            time.Duration
	PrintInterval       time.Duration
	StartFromHead       bool
	LogGroupName        string
	LogStreamNamePrefix string
}

func (w *Watcher) Start() {
	streamNames := []string{}

	for {
		log.Printf("[DEBUG] finding log streams '%s*'", w.LogStreamNamePrefix)

		err := w.CloudWatchLogs.DescribeLogStreamsPages(&cloudwatchlogs.DescribeLogStreamsInput{
			LogGroupName:        aws.String(w.LogGroupName),
			LogStreamNamePrefix: aws.String(w.LogStreamNamePrefix),
		}, func(resp *cloudwatchlogs.DescribeLogStreamsOutput, last bool) bool {
			for _, s := range resp.LogStreams {
				found := false
				for _, ss := range streamNames {
					if ss == *s.LogStreamName {
						found = true
						break
					}
				}

				if !found {
					printer := &Printer{
						CloudWatchLogs: w.CloudWatchLogs,
						Interval:       w.PrintInterval,
						LogGroupName:   w.LogGroupName,
						LogStreamName:  *s.LogStreamName,
						StartFromHead:  w.StartFromHead,
					}
					go printer.Start()

					streamNames = append(streamNames, *s.LogStreamName)
				}
			}
			return true
		})

		if err != nil {
			log.Printf("[ERROR] %s", err)
		}

		time.Sleep(w.Interval)
	}
}
