package documents

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws/awserr"

	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ryotarai/paramedic/awsclient"
)

type Generator struct {
	S3  awsclient.S3
	SSM awsclient.SSM

	ScriptS3Bucket     string
	ScriptS3KeyPrefix  string
	DocumentNamePrefix string
	AgentPath          string
	Region             string
}

func (g *Generator) Create(d *Definition) error {
	scriptKey, err := g.uploadScript(d)
	if err != nil {
		return err
	}

	j, err := g.json(d, scriptKey)
	if err != nil {
		return err
	}

	err = g.createDocument(d, j)
	if err != nil {
		return err
	}

	return nil
}

func (g *Generator) createDocument(d *Definition, content string) error {
	name := fmt.Sprintf("%s%s", g.DocumentNamePrefix, d.Name)

	_, err := g.SSM.DescribeDocument(&ssm.DescribeDocumentInput{
		Name: aws.String(name),
	})
	if err == nil {
		log.Printf("INFO: updating a document '%s'", name)
		resp, err := g.SSM.UpdateDocument(&ssm.UpdateDocumentInput{
			Name:            aws.String(name),
			Content:         aws.String(content),
			DocumentVersion: aws.String("$LATEST"),
		})
		if err != nil {
			return err
		}

		_, err = g.SSM.UpdateDocumentDefaultVersion(&ssm.UpdateDocumentDefaultVersionInput{
			Name:            aws.String(name),
			DocumentVersion: resp.DocumentDescription.DocumentVersion,
		})
		if err != nil {
			return err
		}
	} else {
		aErr, ok := err.(awserr.Error)
		if !ok || aErr.Code() != ssm.ErrCodeInvalidDocument {
			return err
		}
		// document does not exist
		log.Printf("INFO: creating a document '%s'", name)

		_, err := g.SSM.CreateDocument(&ssm.CreateDocumentInput{
			Content:      aws.String(content),
			DocumentType: aws.String("Command"),
			Name:         aws.String(name),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) uploadScript(d *Definition) (string, error) {
	key := fmt.Sprintf("%s%s-%s", g.ScriptS3KeyPrefix, d.Name, d.ScriptSha256())
	log.Printf("INFO: uploading a script to s3://%s/%s", g.ScriptS3Bucket, key)
	input := &s3.PutObjectInput{
		Body:   strings.NewReader(d.Script),
		Bucket: aws.String(g.ScriptS3Bucket),
		Key:    aws.String(key),
	}
	_, err := g.S3.PutObject(input)
	if err != nil {
		return "", err
	}
	return key, nil
}

func (g *Generator) json(d *Definition, scriptKey string) (string, error) {
	j := map[string]interface{}{
		"schemaVersion": "2.2",
		"description":   d.Description,
		"parameters": map[string]map[string]string{
			"executionTimeout": map[string]string{
				"type":           "String",
				"default":        "3600",
				"description":    "(Optional) The time in seconds for a command to be completed before it is considered to have failed. Default is 3600 (1 hour). Maximum is 28800 (8 hours).",
				"allowedPattern": "([1-9][0-9]{0,3})|(1[0-9]{1,4})|(2[0-7][0-9]{1,3})|(28[0-7][0-9]{1,2})|(28800)",
			},
			"outputLogGroup": map[string]string{
				"type":           "String",
				"description":    "(Required) Log group name",
				"allowedPattern": "[\\.\\-_/#A-Za-z0-9]+",
			},
			"outputLogStreamPrefix": map[string]string{
				"type":           "String",
				"description":    "(Required) Log stream name prefix",
				"allowedPattern": "[^:*]*",
			},
			"signalS3Bucket": map[string]string{
				"type":        "String",
				"description": "(Required) S3 bucket the signal object is stored in",
			},
			"signalS3Key": map[string]string{
				"type":        "String",
				"description": "(Required) S3 object key the signal object is stored at",
			},
		},
		"mainSteps": []interface{}{
			map[string]interface{}{
				"action": "aws:runShellScript",
				"name":   "script",
				"inputs": map[string]interface{}{
					"runCommand": []string{
						"export PARAMEDIC_OUTPUT_LOG_GROUP={{outputLogGroup}}",
						"export PARAMEDIC_OUTPUT_LOG_STREAM_PREFIX={{outputLogStreamPrefix}}",
						"export PARAMEDIC_SIGNAL_S3_BUCKET={{signalS3Bucket}}",
						"export PARAMEDIC_SIGNAL_S3_KEY={{signalS3Key}}",
						fmt.Sprintf("export PARAMEDIC_SCRIPT_S3_BUCKET=%s", g.ScriptS3Bucket),
						fmt.Sprintf("export PARAMEDIC_SCRIPT_S3_KEY=%s", scriptKey),
						fmt.Sprintf("export AWS_REGION=%s", g.Region),
						fmt.Sprintf("exec %s", g.AgentPath),
					},
					"timeoutSeconds": "{{ executionTimeout }}",
				},
			},
		},
	}

	b, err := json.Marshal(j)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
