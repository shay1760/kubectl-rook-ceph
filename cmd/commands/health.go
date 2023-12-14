/*
Copyright 2023 The Rook Authors. All rights reserved.

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

package command

import (
	"github.com/rook/kubectl-rook-ceph/pkg/health"
	"github.com/spf13/cobra"
)

var Health = &cobra.Command{
	Use:                "health",
	Short:              "check health of the cluster and common configuration issues",
	DisableFlagParsing: true,
	Args:               cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		verifyOperatorPodIsRunning(cmd.Context(), clientSets)
	},
	Run: func(cmd *cobra.Command, _ []string) {
		health.Health(cmd.Context(), clientSets, operatorNamespace, cephClusterNamespace)
	},
}
