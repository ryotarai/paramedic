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

	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/commands"
	"github.com/ryotarai/paramedic/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uploadCmd represents the upload command
var commandsCancelCmd = &cobra.Command{
	Use:           "cancel",
	Short:         "Cancel a command",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, k := range []string{"command-id", "signal-s3-bucket"} {
			if viper.GetString(k) == "" {
				return fmt.Errorf("%s is required", k)
			}
		}
		commandID := viper.GetString("command-id")
		signalS3Bucket := viper.GetString("signal-s3-bucket")
		signalS3KeyPrefix := viper.GetString("signal-s3-key-prefix")
		signalNo := viper.GetInt("signal")

		awsFactory, err := awsclient.NewFactory()
		if err != nil {
			return err
		}

		err = commands.Cancel(&commands.CancelOptions{
			SSM:   awsFactory.SSM(),
			S3:    awsFactory.S3(),
			Store: store.New(awsFactory.DynamoDB()),

			CommandID:         commandID,
			SignalS3Bucket:    signalS3Bucket,
			SignalS3KeyPrefix: signalS3KeyPrefix,
			SignalNumber:      signalNo,
		})
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	commandsCmd.AddCommand(commandsCancelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	commandsCancelCmd.Flags().String("command-id", "", "Command ID to be canceled")
	commandsCancelCmd.Flags().String("signal-s3-bucket", "", "S3 bucket to store a signal object")
	commandsCancelCmd.Flags().String("signal-s3-key-prefix", "signals/", "S3 key prefix to store a signal object")
	commandsCancelCmd.Flags().Int("signal", 15, "Signal number to be sent to the processes")

	viper.BindPFlags(commandsCancelCmd.Flags())
}
