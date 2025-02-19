/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

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
package api

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use: "api",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
			return
		},
	}

	rootCmd.AddCommand(routerCmd)
	rootCmd.AddCommand(serviceCmd)
	rootCmd.AddCommand(controllerCmd)
	rootCmd.AddCommand(dtoCmd)
	return rootCmd
}
