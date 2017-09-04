// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/google/uuid"
	"github.com/ryotarai/paramedic/awsclient"
	"github.com/spf13/cobra"
)

// uploadCmd represents the upload command
var commandsRunCmd = &cobra.Command{
	Use:           "run",
	Short:         "Run a command",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.New("Hi")
		documentNamePrefix, err := cmd.Flags().GetString("document-name-prefix")
		if err != nil {
			return err
		}
		documentName, err := cmd.Flags().GetString("document-name")
		if err != nil {
			return err
		}
		if documentName == "" {
			return errors.New("document-name is required")
		}
		maxConcurrency, err := cmd.Flags().GetString("max-concurrency")
		if err != nil {
			return err
		}
		maxErrors, err := cmd.Flags().GetString("max-errors")
		if err != nil {
			return err
		}
		instanceIDs, err := cmd.Flags().GetStringSlice("instance-id")
		if err != nil {
			return err
		}
		tags, err := cmd.Flags().GetStringSlice("tag")
		if err != nil {
			return err
		}
		outputLogGroup, err := cmd.Flags().GetString("output-log-group")
		if err != nil {
			return err
		}
		if outputLogGroup == "" {
			return errors.New("output-log-group is required")
		}
		signalS3Bucket, err := cmd.Flags().GetString("signal-s3-bucket")
		if err != nil {
			return err
		}
		if signalS3Bucket == "" {
			return errors.New("signal-s3-bucket is required")
		}
		signalS3KeyPrefix, err := cmd.Flags().GetString("signal-s3-key-prefix")
		if err != nil {
			return err
		}

		documentName = fmt.Sprintf("%s%s", documentNamePrefix, documentName)

		pcommandID := uuid.New().String()
		log.Printf("INFO: pcommand ID is %s", pcommandID)

		awsFactory, err := awsclient.NewFactory()
		if err != nil {
			return err
		}

		targets := []*ssm.Target{}
		if len(instanceIDs) > 0 {
			targets = append(targets, &ssm.Target{
				Key:    aws.String("InstanceIds"),
				Values: aws.StringSlice(instanceIDs),
			})
		}
		for _, t := range tags {
			kv := strings.SplitN(t, ":", 2)
			k := kv[0]
			v := strings.Split(kv[1], ",")

			targets = append(targets, &ssm.Target{
				Key:    aws.String(fmt.Sprintf("tag:%s", k)),
				Values: aws.StringSlice(v),
			})
		}

		ssmClient := awsFactory.SSM()
		resp, err := ssmClient.SendCommand(&ssm.SendCommandInput{
			DocumentName:   aws.String(documentName),
			Targets:        targets,
			MaxConcurrency: aws.String(maxConcurrency),
			MaxErrors:      aws.String(maxErrors),
			Parameters: map[string][]*string{
				"outputLogGroup":        []*string{aws.String(outputLogGroup)},
				"outputLogStreamPrefix": []*string{aws.String(fmt.Sprintf("%s-", pcommandID))},
				"signalS3Bucket":        []*string{aws.String(signalS3Bucket)},
				"signalS3Key":           []*string{aws.String(fmt.Sprintf("%s%s.json", signalS3KeyPrefix, pcommandID))},
			},
		})
		if err != nil {
			return err
		}
		log.Printf("INFO: started a command %s", *resp.Command.CommandId)

		return nil
	},
}

func init() {
	commandsCmd.AddCommand(commandsRunCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	commandsRunCmd.Flags().String("document-name", "", "Document name")
	commandsRunCmd.Flags().String("document-name-prefix", "paramedic-", "Prefix of document name")
	commandsRunCmd.Flags().String("output-log-group", "", "Log group")
	commandsRunCmd.Flags().String("signal-s3-bucket", "", "S3 bucket to store a signal object")
	commandsRunCmd.Flags().String("signal-s3-key-prefix", "signals/", "S3 key prefix to store a signal object")
	commandsRunCmd.Flags().String("max-concurrency", "50", "The maximum number of instances that are allowed to execute the command at the same time")
	commandsRunCmd.Flags().String("max-errors", "50", "The maximum number of errors allowed without the command failing")
	commandsRunCmd.Flags().StringSlice("instance-id", []string{}, "Instance ID")
	commandsRunCmd.Flags().StringSlice("tag", []string{}, "Instance tag")
}
