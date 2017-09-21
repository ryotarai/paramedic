package outputlog

import (
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/ryotarai/paramedic/awsclient"
)

type CloudWatchLogsReader struct {
	CloudWatchLogs  awsclient.CloudWatchLogs
	LogGroup        string
	LogStreamPrefix string
}

func (r *CloudWatchLogsReader) Read() ([]*Event, error) {
	streams, err := r.getLogStreams()
	if err != nil {
		return nil, err
	}
	log.Printf("[DEBUG] %d streams are found: %s", len(streams), streams)

	events := []*Event{}
	for _, s := range streams {
		ev, err := r.getEvents(s)
		if err != nil {
			return nil, err
		}
		events = append(events, ev...)
	}
	SortEventsByTimestamp(events)

	return events, nil
}

func (r *CloudWatchLogsReader) getLogStreams() ([]string, error) {
	streams := []string{}

	err := r.CloudWatchLogs.DescribeLogStreamsPages(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        aws.String(r.LogGroup),
		LogStreamNamePrefix: aws.String(r.LogStreamPrefix),
	}, func(resp *cloudwatchlogs.DescribeLogStreamsOutput, last bool) bool {
		for _, s := range resp.LogStreams {
			streams = append(streams, *s.LogStreamName)
		}
		return true
	})

	if err != nil {
		return nil, err
	}

	return streams, nil
}

func (r *CloudWatchLogsReader) getEvents(logStream string) ([]*Event, error) {
	log.Printf("[DEBUG] Getting log events from %s stream", logStream)
	events := []*Event{}

	input := &cloudwatchlogs.GetLogEventsInput{
		LogGroupName:  aws.String(r.LogGroup),
		LogStreamName: aws.String(logStream),
		StartFromHead: aws.Bool(true),
	}

	err := r.CloudWatchLogs.GetLogEventsPages(input, func(resp *cloudwatchlogs.GetLogEventsOutput, last bool) bool {
		if len(resp.Events) == 0 {
			return false
		}
		for _, e := range resp.Events {
			t := time.Unix(0, (*e.Timestamp)*1000*1000)
			events = append(events, &Event{
				LogStream: logStream,
				Timestamp: t,
				Message:   *e.Message,
			})
		}
		return true
	})
	if err != nil {
		return nil, err
	}

	return events, nil
}
