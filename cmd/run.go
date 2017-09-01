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
	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command",
	Long:  "Run a command",
	RunE: func(cmd *cobra.Command, args []string) error {
		// instanceIds := cmd.Flags().GetStringArray("instance-id")
		// tags := cmd.Flags().GetStringArray("tag")
		// commandFile := args[0]

		// ssm := awsclient.PrepareSSMClient()
		// uploader := uploader.NewUploader(ssm)
		// err := uploader.UploadFile(commandFile)
		// if err != nil {
		// 	return err
		// }

		return nil
	},
	Args: cobra.ExactArgs(1),
}

func init() {
	RootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runCmd.Flags().StringArray("instance-id", []string{}, "Target instance ID")
	runCmd.Flags().StringArray("tag", []string{}, "Target tag (key:value)")
}
