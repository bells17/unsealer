/*
Copyright The KubeVault Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmds

import (
	"kubevault.dev/unsealer/pkg/worker"

	utilerrors "github.com/appscode/go/util/errors"
	v "github.com/appscode/go/version"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"kmodules.xyz/client-go/tools/cli"
)

func NewCmdRun() *cobra.Command {
	opts := worker.NewWorkerOptions()

	cmd := &cobra.Command{
		Use:               "run",
		Short:             "Launch Vault unsealer",
		DisableAutoGenTag: true,
		PreRun: func(c *cobra.Command, args []string) {
			cli.SendPeriodicAnalytics(c, v.Version.Version)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			glog.Infof("Starting operator version %s+%s ...", v.Version.Version, v.Version.CommitHash)

			if errs := opts.Validate(); errs != nil {
				return utilerrors.NewAggregate(errs)
			}
			return opts.Run()
		},
	}

	opts.AddFlags(cmd.Flags())

	return cmd
}
