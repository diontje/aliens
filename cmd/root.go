/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"fmt"
	"os"

	"github.com/diontje/aliens/lib/common"
	"github.com/diontje/aliens/lib/logging"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cmdArgs common.CmdArgs        // commandline args
var log *logrus.Logger            // log handle
var err error

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aliens",
	Short: "aliens is a game simulation",
	Long: `aliens is a game simulations of world being taking over by aliens
from other worlds`,
  // Uncomment the following line if your bare application
  // has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		// nothing to do
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
  if err := rootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func init() {
	cobra.OnInitialize(initialize)
	rootCmd.PersistentFlags().BoolVarP(&cmdArgs.Debug, "debug", "", false,
		"enable debug console logging level")
}

// initialize() process commandline args if any and handle logging setup
func initialize() {
	cmdArgs.Log, err = logging.GetLogger(cmdArgs.Debug)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
