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

	"github.com/ryotarai/paramedic/documents"

	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/ryotarai/paramedic/awsclient"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var documentsListCmd = &cobra.Command{
	Use:           "list",
	Short:         "List documents",
	SilenceUsage:  true,
	SilenceErrors: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		viper.BindPFlags(cmd.Flags())

		awsFactory, err := awsclient.NewFactory()
		if err != nil {
			return err
		}

		ssmClient := awsFactory.SSM()

		err = ssmClient.ListDocumentsPages(&ssm.ListDocumentsInput{}, func(resp *ssm.ListDocumentsOutput, last bool) bool {
			for _, i := range resp.DocumentIdentifiers {
				name := documents.ConvertFromSSMName(*i.Name)
				log.Printf("[INFO] - %s", name)
			}
			return true
		})
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	documentsCmd.AddCommand(documentsListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// uploadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
}
