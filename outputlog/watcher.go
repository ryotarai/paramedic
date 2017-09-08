package outputlog

import (
	"log"
	"sync"
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

	printers []*Printer
}

func (w *Watcher) Follow() {
	for {
		w.findStreams()

		for _, p := range w.printers {
			if !p.Following {
				go p.Follow()
			}
		}

		time.Sleep(w.Interval)
	}
}

func (w *Watcher) Once() {
	w.findStreams()

	var wg sync.WaitGroup
	for _, p := range w.printers {
		wg.Add(1)
		go func(p *Printer) {
			defer wg.Done()
			p.Once()
		}(p)
	}

	wg.Wait()
}

func (w *Watcher) findStreams() {
	log.Printf("[DEBUG] finding log streams '%s*'", w.LogStreamNamePrefix)

	err := w.CloudWatchLogs.DescribeLogStreamsPages(&cloudwatchlogs.DescribeLogStreamsInput{
		LogGroupName:        aws.String(w.LogGroupName),
		LogStreamNamePrefix: aws.String(w.LogStreamNamePrefix),
	}, func(resp *cloudwatchlogs.DescribeLogStreamsOutput, last bool) bool {
		for _, s := range resp.LogStreams {
			found := false
			for _, p := range w.printers {
				if p.LogStreamName == *s.LogStreamName {
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

				w.printers = append(w.printers, printer)
			}
		}
		return true
	})

	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
}
