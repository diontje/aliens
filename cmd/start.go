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
	_"os"
	"strconv"
	"strings"

	"github.com/diontje/aliens/lib/actions"

	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "starts the alien game",
	Long: `start the game of aliens with two aliens if not provided a
n flag specifying the number of alien. For example:

aliens start
  # start the default game with two aliens
aliens start -n 5
  # start the default game with count=5 aliens`,
	Args: func(cmd *cobra.Command, args []string) error {
		cmdArgs.WorldName = strings.TrimSpace(cmdArgs.WorldName)
		cmdArgs.WorldFileLocation = strings.TrimSpace(cmdArgs.WorldFileLocation)
		cmdArgs.NumAliens = 2
		numAliens, _ := cmd.Flags().GetString("numAliens")
		numAliens = strings.TrimSpace(numAliens)
		if len(numAliens) != 0 {
			i, err := strconv.Atoi(numAliens)
			if err != nil {
				return err
			}
			cmdArgs.NumAliens = i
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		log := cmdArgs.Log
		log.Debug("start called")
		err := actions.Start(cmdArgs)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringP("numAliens", "n", "2", "number of aliens")
	startCmd.PersistentFlags().StringVarP(&cmdArgs.WorldName, "world", "", "earth",
		"game file")
	startCmd.PersistentFlags().StringVarP(&cmdArgs.WorldFileLocation, "path", "", "",
		"default game file location (default is ./worlds/)")
}
