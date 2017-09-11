package awsclient

import (
	"github.com/aws/aws-sdk-go/service/ssm"
)

type SSM interface {
	CreateDocument(*ssm.CreateDocumentInput) (*ssm.CreateDocumentOutput, error)
	DescribeDocument(*ssm.DescribeDocumentInput) (*ssm.DescribeDocumentOutput, error)
	UpdateDocument(*ssm.UpdateDocumentInput) (*ssm.UpdateDocumentOutput, error)
	UpdateDocumentDefaultVersion(*ssm.UpdateDocumentDefaultVersionInput) (*ssm.UpdateDocumentDefaultVersionOutput, error)
	SendCommand(*ssm.SendCommandInput) (*ssm.SendCommandOutput, error)
	ListCommands(*ssm.ListCommandsInput) (*ssm.ListCommandsOutput, error)
	ListCommandInvocationsPages(*ssm.ListCommandInvocationsInput, func(*ssm.ListCommandInvocationsOutput, bool) bool) error
	DescribeInstanceInformationPages(*ssm.DescribeInstanceInformationInput, func(*ssm.DescribeInstanceInformationOutput, bool) bool) error
	ListDocumentsPages(*ssm.ListDocumentsInput, func(*ssm.ListDocumentsOutput, bool) bool) error
}
