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

	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/commands"
	"github.com/ryotarai/paramedic/store"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var commandsShowCmd = &cobra.Command{
	Use:           "show",
	Short:         "Show a command",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, k := range []string{"command-id"} {
			if viper.GetString(k) == "" {
				return fmt.Errorf("%s is required", k)
			}
		}
		commandID := viper.GetString("command-id")

		awsFactory, err := awsclient.NewFactory()
		if err != nil {
			return err
		}

		command, err := commands.Get(&commands.GetOptions{
			SSM:       awsFactory.SSM(),
			Store:     store.New(awsFactory.DynamoDB()),
			CommandID: commandID,
		})
		if err != nil {
			return err
		}

		log.Printf("[INFO] Command ID: %s", command.CommandID)
		log.Printf("[INFO] Status: %s", command.Status)
		log.Printf("[INFO] Targets: %s", command.Targets)
		log.Printf("[DEBUG] Paramedic Command ID: %s", command.PcommandID)
		log.Printf("[DEBUG] OutputLogGroup: %s", command.OutputLogGroup)
		log.Printf("[DEBUG] OutputLogStreamPrefix: %s", command.OutputLogStreamPrefix)
		log.Printf("[DEBUG] SignalS3Bucket: %s", command.SignalS3Bucket)
		log.Printf("[DEBUG] SignalS3Key: %s", command.SignalS3Key)

		return nil
	},
}

func init() {
	commandsCmd.AddCommand(commandsShowCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	commandsShowCmd.Flags().String("command-id", "", "Command ID")

	viper.BindPFlags(commandsShowCmd.Flags())
}
