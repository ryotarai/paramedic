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
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ryotarai/paramedic/documents"
	"github.com/ryotarai/paramedic/outputlog"

	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/commands"
	"github.com/ryotarai/paramedic/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uploadCmd represents the upload command
var commandsRunCmd = &cobra.Command{
	Use:           "run",
	Short:         "Run a command",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlags(cmd.Flags())

		for _, k := range []string{"document-name", "signal-s3-bucket"} {
			if viper.GetString(k) == "" {
				return fmt.Errorf("%s is required", k)
			}
		}

		documentName := viper.GetString("document-name")
		maxConcurrency := viper.GetString("max-concurrency")
		maxErrors := viper.GetString("max-errors")
		instanceIDs := viper.GetStringSlice("instance-ids")
		tags := viper.GetStringSlice("tags")
		outputLogGroup := viper.GetString("output-log-group")
		signalS3Bucket := viper.GetString("signal-s3-bucket")
		signalS3KeyPrefix := viper.GetString("signal-s3-key-prefix")

		documentName = documents.ConvertToSSMName(documentName)

		awsFactory, err := awsclient.NewFactory()
		if err != nil {
			return err
		}

		tagMap := map[string][]string{}
		for _, t := range tags {
			parts := strings.SplitN(t, "=", 2)
			tagMap[parts[0]] = []string{parts[1]}
		}

		instances, err := commands.GetInstances(&commands.GetInstancesOptions{
			SSM:         awsFactory.SSM(),
			InstanceIDs: instanceIDs,
			Tags:        tagMap,
		})
		if err != nil {
			return err
		}

		log.Println("[INFO] This command will be executed on the following instances")
		for _, i := range instances {
			log.Printf("[INFO] - %s (%s)", i.ComputerName, i.InstanceID)
		}
		cont, err := askContinue("Are you sure to continue?")
		if err != nil {
			return err
		}
		if !cont {
			fmt.Println("Canceled.")
			return nil
		}
		st := store.New(awsFactory.DynamoDB())

		command, err := commands.Send(&commands.SendOptions{
			SSM:               awsFactory.SSM(),
			Store:             st,
			DocumentName:      documentName,
			InstanceIDs:       instanceIDs,
			Tags:              tagMap,
			MaxConcurrency:    maxConcurrency,
			MaxErrors:         maxErrors,
			OutputLogGroup:    outputLogGroup,
			SignalS3Bucket:    signalS3Bucket,
			SignalS3KeyPrefix: signalS3KeyPrefix,
		})
		if err != nil {
			return err
		}

		log.Printf("[INFO] A command '%s' started", command.CommandID)
		log.Printf("[INFO] To see the command status, run 'paramedic commands show --command-id=%s'", command.CommandID)
		log.Print("[INFO] Output logs will be shown below")

		logStreamPrefix := fmt.Sprintf("%s/", command.PcommandID)
		reader := &outputlog.KinesisReader{
			Kinesis:         awsFactory.Kinesis(),
			StartTimestamp:  time.Now(),
			LogGroup:        outputLogGroup,
			LogStreamPrefix: logStreamPrefix,
		}

		ch := commands.WaitStatus(&commands.WaitStatusOptions{
			SSM:       awsFactory.SSM(),
			Store:     st,
			CommandID: command.CommandID,
			Statuses:  []string{"Success", "Cancelled", "Failed", "TimedOut", "Cancelling"},
		})

		printer := outputlog.NewPrinter(os.Stdout)

		stopCh := make(chan struct{})
		go func() {
			command := <-ch
			log.Printf("[DEBUG] The command is now in %s status.", command.Status)
			time.Sleep(10 * time.Second) // Wait for propagation of logs
			stopCh <- struct{}{}
		}()

		exitCh := make(chan struct{})
		go func() {
			outputlog.Follow(reader, printer, stopCh)
			exitCh <- struct{}{}
		}()

		// Wait until interrupted
		sigCh := make(chan os.Signal)
		signal.Notify(sigCh, syscall.SIGINT)

		select {
		case <-sigCh:
			fmt.Print("Interrupted\n")
			log.Printf("[WARN] The command is NOT cancelled. To cancel, run 'paramedic commands cancel --command-id=%s'", command.CommandID)
		case <-exitCh:
		}

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
	commandsRunCmd.Flags().String("output-log-group", "paramedic", "Log group")
	commandsRunCmd.Flags().String("signal-s3-bucket", "", "S3 bucket to store a signal object")
	commandsRunCmd.Flags().String("signal-s3-key-prefix", "signals/", "S3 key prefix to store a signal object")
	commandsRunCmd.Flags().String("max-concurrency", "50", "The maximum number of instances that are allowed to execute the command at the same time")
	commandsRunCmd.Flags().String("max-errors", "50", "The maximum number of errors allowed without the command failing")
	commandsRunCmd.Flags().StringSlice("instance-ids", []string{}, "Instance IDs")
	commandsRunCmd.Flags().StringSlice("tags", []string{}, "Instance tags (e.g. 'Role=app,Env=prod')")
}
