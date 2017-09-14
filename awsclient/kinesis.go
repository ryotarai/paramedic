package awsclient

import "github.com/aws/aws-sdk-go/service/kinesis"

type Kinesis interface {
	GetRecords(*kinesis.GetRecordsInput) (*kinesis.GetRecordsOutput, error)
	GetShardIterator(*kinesis.GetShardIteratorInput) (*kinesis.GetShardIteratorOutput, error)
	DescribeStream(*kinesis.DescribeStreamInput) (*kinesis.DescribeStreamOutput, error)
}
