package outputlog

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kinesis"

	"github.com/ryotarai/paramedic/awsclient"
)

const kinesisStreamName = "paramedic-logs" // TODO: configurable

type KinesisReader struct {
	Kinesis         awsclient.Kinesis
	StartTimestamp  time.Time
	LogGroup        string
	LogStreamPrefix string

	shardIterator map[string]string // map[shard ID]iterator
}

type kinesisRecord struct {
	MessageType         string            `json:"messageType"`
	Owner               string            `json:"owner"`
	LogGroup            string            `json:"logGroup"`
	LogStream           string            `json:"logStream"`
	SubscriptionFilters []string          `json:"subscriptionFilters"`
	LogEvents           []kinesisLogEvent `json:"logEvents"`
}
type kinesisLogEvent struct {
	ID        string `json:"id"`
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}

func (r *KinesisReader) Read() ([]*Event, error) {
	if r.shardIterator == nil {
		err := r.initShardIterator()
		if err != nil {
			return nil, err
		}
	}

	events := []*Event{}
	for shardID, iter := range r.shardIterator {
		log.Printf("[DEBUG] Getting records from Kinesis Streams (shard ID: %s, iterator: %s)", shardID, iter)
		resp, err := r.Kinesis.GetRecords(&kinesis.GetRecordsInput{
			ShardIterator: aws.String(iter),
		})
		if err != nil {
			return nil, err
		}
		r.shardIterator[shardID] = *resp.NextShardIterator

		ev, err := r.recordsToEvents(resp.Records)
		if err != nil {
			return nil, err
		}
		events = append(events, ev...)
	}

	SortEventsByTimestamp(events)

	return events, nil
}

func (r *KinesisReader) recordsToEvents(records []*kinesis.Record) ([]*Event, error) {
	events := []*Event{}
	for _, r1 := range records {
		reader, err := gzip.NewReader(bytes.NewReader(r1.Data))
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(reader)
		if err != nil {
			return nil, err
		}

		r2 := kinesisRecord{}
		err = json.Unmarshal(b, &r2)
		if err != nil {
			return nil, err
		}

		if r2.MessageType != "DATA_MESSAGE" {
			continue // ignore
		}
		if r2.LogGroup != r.LogGroup || !strings.HasPrefix(r2.LogStream, r.LogStreamPrefix) {
			continue // ignore
		}

		for _, e := range r2.LogEvents {
			events = append(events, &Event{
				Message:   e.Message,
				Timestamp: time.Unix(0, e.Timestamp*1000*1000),
				LogStream: r2.LogStream,
			})
		}
	}
	return events, nil
}

func (r *KinesisReader) initShardIterator() error {
	log.Printf("[DEBUG] Getting initial shart iterator")
	r.shardIterator = map[string]string{}

	resp, err := r.Kinesis.DescribeStream(&kinesis.DescribeStreamInput{
		StreamName: aws.String(kinesisStreamName),
	})
	if err != nil {
		return err
	}

	for _, s := range resp.StreamDescription.Shards {
		id := *s.ShardId
		resp, err := r.Kinesis.GetShardIterator(&kinesis.GetShardIteratorInput{
			StreamName:        aws.String(kinesisStreamName),
			ShardId:           aws.String(id),
			ShardIteratorType: aws.String("AT_TIMESTAMP"),
			Timestamp:         aws.Time(r.StartTimestamp),
		})
		if err != nil {
			return err
		}
		r.shardIterator[id] = *resp.ShardIterator
	}
	return nil
}
