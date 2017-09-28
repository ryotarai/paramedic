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
	"log"

	"github.com/ryotarai/paramedic/awsclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uploadCmd represents the upload command
var commandsCancelCmd = &cobra.Command{
	Use:           "cancel",
	Short:         "Cancel a command",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE:          commandsCancelHandler,
}

func commandsCancelHandler(cmd *cobra.Command, args []string) error {
	viper.BindPFlags(cmd.Flags())

	if err := requireStringFlags([]string{"command-id"}); err != nil {
		return err
	}

	commandID := viper.GetString("command-id")
	signalNo := viper.GetInt("signal")

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

	err = cmdClient.Cancel(command, signalNo)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Canceling a command %s", commandID)

	canceled := []string{"Success", "Cancelled", "Failed", "TimedOut", "Cancelling"}
	command = <-cmdClient.WaitStatus(command.CommandID, canceled)

	log.Printf("[INFO] The command is now in %s state", command.Status)

	return nil
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
	commandsCancelCmd.Flags().Int("signal", 15, "Signal number to be sent to the processes")
}
