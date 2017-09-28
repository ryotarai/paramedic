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
	"log"
	"os"
	"time"

	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/outputlog"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uploadCmd represents the upload command
var commandsLogCmd = &cobra.Command{
	Use:           "log",
	Short:         "Show logs of a command",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE:          commandsLogHandler,
}

func commandsLogHandler(cmd *cobra.Command, args []string) error {
	viper.BindPFlags(cmd.Flags())

	for _, k := range []string{"command-id", "output-log-group"} {
		if viper.GetString(k) == "" {
			return fmt.Errorf("%s is required", k)
		}
	}
	outputLogGroup := viper.GetString("output-log-group")
	commandID := viper.GetString("command-id")
	follow := viper.GetBool("follow")

	awsf, err := awsclient.NewFactory()
	if err != nil {
		return err
	}

	cmdClient, err := newCommandsClient(awsf)
	if err != nil {
		return err
	}

	command, err := cmdClient.Get(commandID)
	if err != nil {
		return err
	}

	logStreamPrefix := fmt.Sprintf("%s/", command.PcommandID)

	var reader outputlog.Reader
	if follow {
		reader = &outputlog.KinesisReader{
			Kinesis:         awsf.Kinesis(),
			StartTimestamp:  time.Now(),
			LogGroup:        outputLogGroup,
			LogStreamPrefix: logStreamPrefix,
		}
	} else {
		reader = &outputlog.CloudWatchLogsReader{
			CloudWatchLogs:  awsf.CloudWatchLogs(),
			LogGroup:        outputLogGroup,
			LogStreamPrefix: logStreamPrefix,
		}
	}

	printer := outputlog.NewPrinter(os.Stdout)

	if follow {
		stopCh := make(chan struct{})
		go func() {
			command := <-cmdClient.WaitStatus(commandID, []string{"Success", "Cancelled", "Failed", "TimedOut", "Cancelling"})
			log.Printf("[DEBUG] The command is now in %s status.", command.Status)
			time.Sleep(10 * time.Second) // Wait for propagation of logs
			stopCh <- struct{}{}
		}()
		outputlog.Follow(reader, printer, stopCh)
	} else {
		events, err := reader.Read()
		if err != nil {
			return err
		}
		printer.Print(events)
	}

	return nil
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
	commandsLogCmd.Flags().BoolP("follow", "f", false, "Follow logs like `tail -f -n0` (Kinesis Streams will be used)")
}
