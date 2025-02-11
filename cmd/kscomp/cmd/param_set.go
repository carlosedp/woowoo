// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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
	"github.com/bryanl/woowoo/action"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	vParamSetIndex = "param-set-index"
)

// setCmd represents the set command
var paramSetCmd = &cobra.Command{
	Use:   "set",
	Short: "param set",
	Long:  `param set`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 3 {
			logrus.Fatal("set <component-name> <param-key> <param-value>")
		}

		indexOpt := action.ParamSetWithIndex(viper.GetInt(vParamSetIndex))
		return action.ParamSet(fs, args[0], args[1], args[2], indexOpt)
	},
}

func init() {
	paramCmd.AddCommand(paramSetCmd)

	paramSetCmd.Flags().IntP(flagIndex, "i", 0, "Index in manifest")
	viper.BindPFlag(vParamSetIndex, paramSetCmd.Flags().Lookup(flagIndex))

	// TODO: support env
}
