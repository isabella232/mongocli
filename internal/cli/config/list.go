// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"

	"github.com/mongodb/mongocli/internal/config"
	"github.com/spf13/cobra"
)

func ListBuilder() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   listShort,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Available profiles:")
			profiles := config.List()
			for _, p := range profiles {
				fmt.Printf("  %s\n", p)
			}
		},
	}

	return cmd
}
