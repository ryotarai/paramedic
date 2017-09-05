// Copyright © 2017 Ryota Arai <ryota.arai@gmail.com>
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
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uploadCmd represents the upload command
var commandsLogCmd = &cobra.Command{
	Use:           "log",
	Short:         "Show logs of a command",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, k := range []string{"command-id", "output-log-group"} {
			if viper.GetString(k) == "" {
				return fmt.Errorf("%s is required", k)
			}
		}
		outputLogGroup := viper.GetString("output-log-group")
		commandID := viper.GetString("command-id")
		fromHead := viper.GetBool("from-head")

		awsFactory, err := awsclient.NewFactory()
		if err != nil {
			return err
		}

		st := store.New(awsFactory.DynamoDB())
		r, err := st.GetCommand(commandID)
		if err != nil {
			return err
		}
		pcommandID := r.PcommandID

		cwlogsClient := awsFactory.CloudWatchLogs()

		streamNames := []string{}
		err = cwlogsClient.DescribeLogStreamsPages(&cloudwatchlogs.DescribeLogStreamsInput{
			LogGroupName:        aws.String(outputLogGroup),
			LogStreamNamePrefix: aws.String(fmt.Sprintf("%s/", pcommandID)),
		}, func(resp *cloudwatchlogs.DescribeLogStreamsOutput, last bool) bool {
			for _, s := range resp.LogStreams {
				streamNames = append(streamNames, *s.LogStreamName)
			}
			return true
		})
		if err != nil {
			return err
		}

		for _, streamName := range streamNames {
			parts := strings.Split(streamName, "/")
			instanceID := parts[len(parts)-1]
			err = cwlogsClient.GetLogEventsPages(&cloudwatchlogs.GetLogEventsInput{
				LogGroupName:  aws.String(outputLogGroup),
				LogStreamName: aws.String(streamName),
				StartFromHead: aws.Bool(fromHead),
			}, func(resp *cloudwatchlogs.GetLogEventsOutput, last bool) bool {
				if len(resp.Events) == 0 {
					return false
				}
				for _, e := range resp.Events {
					t := time.Unix(0, (*e.Timestamp)*1000*1000)
					fmt.Printf("%s | %s | %s\n", instanceID, t.Format(time.RFC3339), *e.Message)
				}
				return true
			})
		}

		return nil
	},
}

func init() {
	commandsCmd.AddCommand(commandsLogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	commandsLogCmd.Flags().String("command-id", "", "Command ID")
	commandsLogCmd.Flags().String("output-log-group", "", "Log group")
	commandsLogCmd.Flags().Bool("from-head", false, "Read logs from head")

	viper.BindPFlags(commandsLogCmd.Flags())
}
