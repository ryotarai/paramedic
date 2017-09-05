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
	"errors"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/ssm"

	"github.com/ryotarai/paramedic/awsclient"
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

		ssmClient := awsFactory.SSM()
		resp, err := ssmClient.ListCommands(&ssm.ListCommandsInput{
			CommandId: aws.String(commandID),
		})
		if err != nil {
			return err
		}

		if len(resp.Commands) == 0 {
			return errors.New("command is not found")
		}
		status := *resp.Commands[0].Status
		if status != "Pending" && status != "InProgress" {
			return fmt.Errorf("can't cancel the command because its status is %s", status)
		}

		// Get pcommand ID
		st := store.New(awsFactory.DynamoDB())
		r, err := st.GetID(commandID)
		if err != nil {
			return err
		}

		log.Printf("DEBUG: PcommandID is %s", r.PcommandID)

		// Put signal

		// Wait for command

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

	viper.BindPFlags(commandsCancelCmd.Flags())
}
