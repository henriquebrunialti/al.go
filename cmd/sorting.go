/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// sortingCmd represents the sorting command
var sortingCmd = &cobra.Command{
	Use:   "sorting",
	Short: "Sorting Algorithms",
	Long: `These are the sortings algorithms`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("These are some sorting algorithms animations")
		fmt.Printf("\nUse -h to list all supported animations.\n")

		fmt.Printf("\nANIMATIONS:\n")
		fmt.Printf("\nal.go sorting mergesort		Visualization of the MergeSort algorithm")
	},
}

func init() {
	rootCmd.AddCommand(sortingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sortingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sortingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
