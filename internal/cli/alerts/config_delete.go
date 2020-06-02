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

package alerts

import (
	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/description"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

type ConfigDeleteOpts struct {
	cli.GlobalOpts
	*cli.DeleteOpts
	store store.AlertConfigurationDeleter
}

func (opts *ConfigDeleteOpts) initStore() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *ConfigDeleteOpts) Run() error {
	return opts.Delete(opts.store.DeleteAlertConfiguration, opts.ConfigProjectID())
}

// mongocli atlas alerts config(s) delete <ID> --projectId projectId [--confirm]
func ConfigDeleteBuilder() *cobra.Command {
	opts := &ConfigDeleteOpts{
		DeleteOpts: cli.NewDeleteOpts("Alert config '%s' deleted\n", "Alert config not deleted"),
	}

	cmd := &cobra.Command{
		Use:     "delete <ID>",
		Short:   description.DeleteAlertsConfig,
		Aliases: []string{"rm", "Delete", "Remove"},
		Args:    cobra.ExactArgs(1),
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(opts.initStore); err != nil {
				return err
			}
			opts.Entry = args[0]
			return opts.Prompt()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().BoolVar(&opts.Confirm, flag.Force, false, usage.Force)

	cmd.Flags().StringVar(&opts.ProjectID, flag.ProjectID, "", usage.ProjectID)

	return cmd
}