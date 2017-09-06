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
	"time"

	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/outputlog"
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

		watcher := &outputlog.Watcher{
			CloudWatchLogs:      awsFactory.CloudWatchLogs(),
			Interval:            30 * time.Second,
			PrintInterval:       30 * time.Second,
			StartFromHead:       fromHead,
			LogGroupName:        outputLogGroup,
			LogStreamNamePrefix: fmt.Sprintf("%s/", pcommandID),
		}
		watcher.Start()

		// TODO: exit if the command finished

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
