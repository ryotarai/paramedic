package documents

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/ryotarai/paramedic/awsclient"
)

type Client struct {
	SSM awsclient.SSM
	S3  awsclient.S3

	ScriptS3Bucket    string
	ScriptS3KeyPrefix string
}

func (c *Client) Create(d *Definition) error {
	err := c.uploadScript(d)
	if err != nil {
		return err
	}

	name := ConvertToSSMName(d.Name)
	content, err := d.DocumentContent(c.ScriptS3Bucket, c.scriptKey(d))
	if err != nil {
		return err
	}

	_, err = c.SSM.DescribeDocument(&ssm.DescribeDocumentInput{
		Name: aws.String(name),
	})
	if err == nil {
		// existing
		c.updateDocument(name, content)
	} else {
		aErr, ok := err.(awserr.Error)
		if !ok || aErr.Code() != ssm.ErrCodeInvalidDocument {
			return err
		}
		// document does not exist
		c.createDocument(name, content)
	}
	return nil
}

func (c *Client) scriptKey(d *Definition) string {
	return fmt.Sprintf("%s%s-%s", c.ScriptS3KeyPrefix, d.Name, d.ScriptSha256())
}

func (c *Client) uploadScript(d *Definition) error {
	key := c.scriptKey(d)
	log.Printf("[INFO] Uploading a script to s3://%s/%s", c.ScriptS3Bucket, key)
	input := &s3.PutObjectInput{
		Body:   strings.NewReader(d.Script),
		Bucket: aws.String(c.ScriptS3Bucket),
		Key:    aws.String(key),
	}
	_, err := c.S3.PutObject(input)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) updateDocument(name, content string) error {
	log.Printf("[INFO] Updating a document '%s'", name)
	resp, err := c.SSM.UpdateDocument(&ssm.UpdateDocumentInput{
		Name:            aws.String(name),
		Content:         aws.String(content),
		DocumentVersion: aws.String("$LATEST"),
	})
	if err != nil {
		return err
	}

	_, err = c.SSM.UpdateDocumentDefaultVersion(&ssm.UpdateDocumentDefaultVersionInput{
		Name:            aws.String(name),
		DocumentVersion: resp.DocumentDescription.DocumentVersion,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) createDocument(name, content string) error {
	log.Printf("[INFO] Creating a document '%s'", name)

	_, err := c.SSM.CreateDocument(&ssm.CreateDocumentInput{
		Content:      aws.String(content),
		DocumentType: aws.String("Command"),
		Name:         aws.String(name),
	})
	if err != nil {
		return err
	}
	return nil
}
