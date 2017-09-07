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
	"github.com/ryotarai/paramedic/awsclient"
	"github.com/ryotarai/paramedic/documents"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// uploadCmd represents the upload command
var uploadCmd = &cobra.Command{
	Use:           "upload",
	Short:         "Upload a document",
	Args:          cobra.ExactArgs(1),
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlags(cmd.Flags())

		scriptS3Bucket := viper.GetString("script-s3-bucket")
		scriptS3KeyPrefix := viper.GetString("script-s3-key-prefix")
		documentNamePrefix := viper.GetString("document-name-prefix")
		agentPath := viper.GetString("agent-path")

		awsFactory, err := awsclient.NewFactory()
		if err != nil {
			return err
		}

		def, err := documents.LoadDefinition(args[0])
		if err != nil {
			return err
		}

		generator := documents.Generator{
			S3:                 awsFactory.S3(),
			SSM:                awsFactory.SSM(),
			ScriptS3Bucket:     scriptS3Bucket,
			ScriptS3KeyPrefix:  scriptS3KeyPrefix,
			DocumentNamePrefix: documentNamePrefix,
			AgentPath:          agentPath,
			Region:             awsFactory.Region(),
		}
		err = generator.Create(def)
		if err != nil {
			return err
		}

		// awsFactory := awsclient.NewFactory()
		// ssm := awsFactory.SSM()

		// uploader := documents.NewUploader()

		return nil
	},
}

func init() {
	documentsCmd.AddCommand(uploadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	uploadCmd.Flags().String("script-s3-bucket", "", "S3 bucket to store a script file")
	uploadCmd.Flags().String("script-s3-key-prefix", "", "S3 key prefix to store a script file")
	uploadCmd.Flags().String("document-name-prefix", "paramedic-", "Prefix of document name")
	uploadCmd.Flags().String("agent-path", "paramedic-agent", "Path to paramedic-agent binary")
}
