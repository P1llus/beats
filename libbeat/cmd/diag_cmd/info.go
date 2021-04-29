// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package diag_cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/cmd/instance"
	"github.com/elastic/beats/v7/libbeat/diagnostics"
)

func GenDiagInfoCmd(settings instance.Settings, beatCreator beat.Creator) *cobra.Command {
	genTemplateConfigCmd := &cobra.Command{
		Use:   "info",
		Short: "Export defined dashboard to stdout",
		Run: func(cmd *cobra.Command, args []string) {
			b, err := instance.NewInitializedBeat(settings)
			if err != nil {
				fmt.Println("failed to initialize beat")
			}
			var config map[string]interface{}
			err = b.RawConfig.Unpack(&config)
			if err != nil {
				fmt.Println("Error unpacking config")
			}

			go diagnostics.GetInfo(b, config)
			err = instance.Run(settings, beatCreator)
			if err != nil {
				os.Exit(1)
			}
		},
	}
	return genTemplateConfigCmd
}
